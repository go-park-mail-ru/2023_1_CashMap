package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	utildb "depeche/internal/repository/utils"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostStorage struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) repository.PostRepository {
	return &PostStorage{
		db: db,
	}
}

func (storage *PostStorage) GetPostSenderInfo(postID uint) (*entities.UserInfo, *entities.CommunityInfo, error) {
	author := &entities.UserInfo{}
	err := storage.db.Get(author, PostSenderInfoQuery, postID)
	if err == sql.ErrNoRows {
		return nil, nil, apperror.NewServerError(apperror.PostNotFound, fmt.Errorf("post with id=%d not found", postID))
	}

	if err != nil {
		return nil, nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	// TODO: Может не работать
	community := &entities.CommunityInfo{}
	err = storage.db.Get(community, CommunityPostInfoQuery, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			community = nil
		} else {
			return nil, nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
	}

	return author, community, nil
}

func (storage *PostStorage) SelectPostById(postId uint, email string) (*entities.Post, error) {
	post := &entities.Post{}

	err := storage.db.Get(post, PostInfoByIdQuery, postId, email)
	if err == sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.PostNotFound, fmt.Errorf("post with id=%d not found", postId))
	}
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return post, nil
}

func (storage *PostStorage) SelectPostsByCommunityLink(info *dto.PostsGetByLink, email string) ([]*entities.Post, error) {
	// больше тот, кто запощен позже
	//TODO: NamedQueryx
	rows, err := storage.db.Queryx(
		PostByCommunityLinkQuery,
		*info.CommunityLink,
		info.LastPostDate,
		info.BatchSize,
		email)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	posts, err := getSliceFromRows[entities.Post](rows, info.BatchSize)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return posts, nil
}

func (storage *PostStorage) SelectPostsByUserLink(info *dto.PostsGetByLink, email string) ([]*entities.Post, error) {
	rows, err := storage.db.Queryx(
		PostsByUserLinkQuery,
		info.OwnerLink,
		info.LastPostDate,
		info.BatchSize,
		email)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	posts, err := getSliceFromRows[entities.Post](rows, info.BatchSize)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return posts, nil
}

func (storage *PostStorage) CheckReadAccess(senderEmail string) (bool, error) {
	//TODO: сделать проверку поля privacy (если private - проверить, что подписан) - РК3
	// TODO: НА ДАЛЕКОЕ БУДУЩЕЕ(если вообще докрутим) - ПРОВЕРКА НА ВОЗМОЖНОСТЬ ДОСТУПА ЮЗЕРА К ПРОФИЛЮ ДРУОГО ЮЗЕРА
	return true, nil
}

func (storage *PostStorage) CheckWriteAccess(senderEmail string, dto *dto.PostCreate) (bool, error) {
	//TODO: сделать проверку поля privacy (если private - проверить, что подписан) - РК3
	if dto.CommunityLink != nil {
		//TODO: сделать проверку поля privacy (если не open - проверить, что имеет права через Community management) - РК3
	} else {
		var accessType string
		err := storage.db.Get(&accessType, "SELECT access_to_posts FROM UserProfile WHERE link = $1", dto.OwnerLink)
		if err == sql.ErrNoRows {
			return false, apperror.NewServerError(apperror.UserNotFound, fmt.Errorf("user with link=%s not found", *dto.OwnerLink))
		}
		if err != nil {
			return false, apperror.NewServerError(apperror.InternalServerError, err)
		}

		switch accessType {
		case "all":
			return true, nil
		case "friends":
			//TODO: проверка на друзей.....
			return true, nil
		default:
			return false, apperror.NewServerError(apperror.InternalServerError, fmt.Errorf("error in db record: invalid access type \"%s\"", accessType))
		}
	}

	return true, nil
}

func (storage *PostStorage) CreatePost(senderEmail string, dto *dto.PostCreate) (uint, error) {
	currentTime := utils.CurrentTimeString()

	tx, err := storage.db.Beginx()
	if err != nil {
		return 0, err
	}

	var communityID *uint
	var ownerID *uint
	if dto.CommunityLink != nil {
		communityID = new(uint)
		err = tx.Get(communityID, "SELECT id FROM groups WHERE link = $1", *dto.CommunityLink)
		if err == sql.ErrNoRows {
			// Неизвестый link сообщества
			if err := tx.Rollback(); err != nil {
				return 0, apperror.NewServerError(apperror.InternalServerError, err)
			}
			return 0, apperror.NewServerError(apperror.CommunityNotFound, fmt.Errorf("community with link \"%s\" not found", *dto.CommunityLink))
		}
		if err != nil {
			return 0, apperror.NewServerError(apperror.InternalServerError, err)
		}
	} else if dto.OwnerLink != nil {
		ownerID = new(uint)
		err = tx.Get(ownerID, "SELECT id FROM UserProfile WHERE link = $1", *dto.OwnerLink)
		if err == sql.ErrNoRows && dto.OwnerLink != nil && dto.CommunityLink == nil {
			// Неизвестый link юзера
			if err := tx.Rollback(); err != nil {
				return 0, err
			}
			return 0, apperror.NewServerError(apperror.UserNotFound, fmt.Errorf("user with link \"%s\" not found", *dto.OwnerLink))
		}
		if err != nil {
			return 0, apperror.NewServerError(apperror.InternalServerError, err)
		}

	}

	query, err := tx.Prepare(CreatePostQuery)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, apperror.NewServerError(apperror.InternalServerError, err)
		}
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}

	var postID uint
	rows, err := query.Query(communityID, senderEmail, ownerID, dto.ShouldShowAuthor, dto.Text, currentTime, currentTime)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err == sql.ErrNoRows {
		return 0, apperror.NewServerError(apperror.InternalServerError, errors.New("just created post not found"))
	}
	if err != nil {
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}

	if rows.Next() {
		rows.Scan(&postID)
	}
	if err != nil {
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}

	// TODO: может выстрелить в ногу
	return postID, nil
}

func (storage *PostStorage) UpdatePost(senderEmail string, dto *dto.PostUpdate) error {
	var isAuthor bool
	tx, err := storage.db.Beginx()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	err = tx.Get(&isAuthor, "SELECT true FROM UserProfile AS profile JOIN Post as post ON profile.id = post.author_id WHERE post.id = $1 AND email = $2", dto.PostID, senderEmail)
	if err == sql.ErrNoRows {
		return apperror.NewServerError(apperror.PostNotFound, fmt.Errorf("post with id=%d and author=%s not found", *dto.PostID, senderEmail))
	}
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	if !isAuthor {
		return apperror.NewServerError(apperror.PostEditingNowAllowed, fmt.Errorf("post editing with id=%d for author=%s is not allowed", *dto.PostID, senderEmail))
	}

	dtoToDB := map[string]string{
		"Text":             "text_content",
		"ShouldShowAuthor": "show_author",
		"ChangeDate":       "change_date",
	}
	dto.ChangeDate = utils.CurrentTimeString()
	err = repository.UpdateTable(storage.db, "Post", "id", *dto.PostID, dtoToDB, dto)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil
}

func (storage *PostStorage) DeletePost(senderEmail string, dto *dto.PostDelete) error {
	result, err := storage.db.Exec("UPDATE Post SET is_deleted = true WHERE author_id = (SELECT id FROM UserProfile WHERE email = $1) AND id = $2",
		senderEmail,
		dto.PostID)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	deletedCount, err := result.RowsAffected()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	if deletedCount == 0 {
		return apperror.NewServerError(apperror.PostEditingNowAllowed, errors.New("resource doesn't exists or delete isn't allowed"))
	}

	return nil
}

func (storage *PostStorage) SetLike(email string, postID uint) error {
	// TODO: нужна проверка доступа к постам (ручка недописанная выше)
	tx, err := storage.db.Beginx()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	_, err = tx.Exec(SetLikeQuery, postID, email)
	if err != nil {
		// TODO: выделить отдельно ошибку при нарушении констрэнта unique
		_ = tx.Rollback()
		return apperror.NewServerError(apperror.AlreadyLiked, err)
	}

	_, err = tx.Exec("UPDATE Post SET likes_amount = likes_amount + 1 WHERE id = $1", postID)
	if err != nil {
		_ = tx.Rollback()
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	err = tx.Commit()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil
}

func (storage *PostStorage) CancelLike(email string, postID uint) error {
	// TODO: нужна проверка доступа к постам (ручка недописанная выше)
	tx, err := storage.db.Beginx()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	execResult, err := tx.Exec(CancelLikeQuery, postID, email)
	if err != nil {
		// TODO: выделить отдельно ошибку при нарушении констрэнта unique
		_ = tx.Rollback()
		return apperror.NewServerError(apperror.AlreadyLiked, err)
	}

	rowsAmount, err := execResult.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	if rowsAmount == 0 {
		_ = tx.Rollback()
		return apperror.NewServerError(apperror.LikeIsMissing,
			fmt.Errorf("like for post with id = %d for user with email = %s doesn't exists", postID, email))
	}

	rows, err := tx.Queryx("UPDATE Post SET likes_amount = likes_amount - 1 WHERE id = $1", postID)
	defer utildb.CloseRows(rows)
	if err != nil {
		_ = tx.Rollback()
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	err = tx.Commit()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil
}

func (storage *PostStorage) GetLikesAmount(email string, postID uint) (int, error) {
	// TODO: нужна проверка доступа к постам (ручка недописанная выше)
	var likesAmount int
	err := storage.db.Get(&likesAmount, "SELECT likes_amount FROM Post WHERE id = $1", postID)
	if err != nil {
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return likesAmount, nil
}

func getSliceFromRows[T any](rows *sqlx.Rows, size uint) ([]*T, error) {
	items := make([]*T, 0, size)
	it := 0
	for rows.Next() {
		item := new(T)
		err := rows.StructScan(item)
		if err != nil {
			return nil, err
		}
		it++
		items = append(items, item)
	}
	return items, nil
}
