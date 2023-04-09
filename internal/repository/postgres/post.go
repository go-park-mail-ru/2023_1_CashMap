package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/utils"
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
	owner := &entities.UserInfo{}
	err := storage.db.Get(owner, "SELECT first_name, last_name, url, link FROM Post as post"+
		" JOIN UserProfile as profile ON post.author_id = profile.id"+
		" LEFT JOIN Photo as photo ON profile.avatar_id = photo.id WHERE post.id = $1", postID)
	if err != nil {
		return nil, nil, err
	}

	// TODO: Может не работать
	community := &entities.CommunityInfo{}
	err = storage.db.Get(community, "SELECT title, url, link FROM Post as post"+
		" JOIN Community as community ON post.community_id = community.id"+
		" LEFT JOIN Photo as photo ON community.avatar_id = photo.id WHERE post.id = $1", postID)
	if err != nil {
		if err == sql.ErrNoRows {
			community = nil
		} else {
			return nil, nil, err
		}
	}

	return owner, community, nil
}

func (storage *PostStorage) SelectPostById(postId uint) (*entities.Post, error) {
	post := &entities.Post{}

	err := storage.db.Get(post, "SELECT post.id, text_content, author.link as author_link, post.likes_amount, post.show_author, post.creation_date"+
		" FROM Post AS post JOIN UserProfile AS author ON post.author_id = author.id"+
		" LEFT JOIN Community as community on post.community_id = community.id"+
		" LEFT JOIN UserProfile as owner ON post.owner_id = owner.id"+
		" WHERE post.id = $1 AND post.is_deleted = false", postId)
	if err == sql.ErrNoRows {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (storage *PostStorage) SelectPostsByCommunityLink(info *dto.PostsGetByLink) ([]*entities.Post, error) {
	// больше тот, кто запощен позже
	//TODO: NamedQueryx
	rows, err := storage.db.Queryx("SELECT post.id, text_content, author.link as author_link, post.likes_amount, post.show_author, post.creation_date, post.change_date "+
		"FROM Post AS post JOIN UserProfile AS author ON post.author_id = author.id "+
		"LEFT JOIN Community as community on post.community_id = community.id "+
		"WHERE post.community_id = (SELECT id FROM Community WHERE link = $1) AND post.creation_date > $2 AND post.is_deleted = false ORDER BY post.creation_date DESC LIMIT $3",
		*info.CommunityLink,
		info.LastPostDate,
		info.BatchSize)
	if err != nil {
		return nil, err
	}

	posts, err := getSliceFromRows[entities.Post](rows, info.BatchSize)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (storage *PostStorage) SelectPostsByUserLink(info *dto.PostsGetByLink) ([]*entities.Post, error) {
	rows, err := storage.db.Queryx("SELECT post.id, text_content, author.link as author_link, post.likes_amount, post.show_author, post.creation_date, post.change_date "+
		"FROM Post AS post JOIN UserProfile AS author ON post.author_id = author.id "+
		"LEFT JOIN Community as community on post.community_id = community.id "+
		"LEFT JOIN UserProfile as owner ON post.owner_id = owner.id "+
		"WHERE post.owner_id = (SELECT id FROM UserProfile WHERE link = $1) AND post.creation_date > $2 AND post.is_deleted = false ORDER BY post.creation_date DESC LIMIT $3",
		info.OwnerLink,
		info.LastPostDate,
		info.BatchSize)
	if err != nil {
		return nil, err
	}

	posts, err := getSliceFromRows[entities.Post](rows, info.BatchSize)
	if err != nil {
		return nil, err
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
			return false, errors.New("not found")
		}
		if err != nil {
			return false, err
		}

		switch accessType {
		case "all":
			return true, nil
		case "friends":
			//TODO: проверка на друзей.....
			return true, nil
		default:
			return false, errors.New("error in db record")
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
		err = tx.Get(communityID, "SELECT id FROM Community WHERE link = $1", *dto.CommunityLink)
		if err != nil {
			return 0, err
		}
		// Неизвестый link сообщества
		if communityID == nil {
			if err := tx.Rollback(); err != nil {
				return 0, err
			}
			return 0, errors.New("unknown community link")
		}
	} else if dto.OwnerLink != nil {
		ownerID = new(uint)
		err = tx.Get(ownerID, "SELECT id FROM UserProfile WHERE link = $1", *dto.OwnerLink)
		if err != nil {
			return 0, err
		}
		// Неизвестый link юзера
		if dto.OwnerLink != nil && dto.CommunityLink == nil && ownerID == nil {
			if err := tx.Rollback(); err != nil {
				return 0, err
			}
			return 0, errors.New("unknown profile link")
		}
	}

	query, err := tx.PrepareNamed("INSERT INTO Post (community_id, author_id, owner_id, show_author, text_content, creation_date, change_date) " +
		"VALUES (:community_id, (SELECT id FROM UserProfile WHERE email = :sender_email), :owner_id, :show_author, :text, :init_time, :change_time) RETURNING id")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	var postID uint
	err = query.Get(&postID, map[string]interface{}{
		"owner_id":     ownerID,
		"sender_email": senderEmail,
		"community_id": communityID,
		"show_author":  dto.ShouldShowAuthor,
		"text":         dto.Text,
		"init_time":    currentTime,
		"change_time":  currentTime,
	})
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	// TODO: может выстрелить в ногу
	return postID, nil
}

func (storage *PostStorage) UpdatePost(senderEmail string, dto *dto.PostUpdate) error {
	var isAuthor bool
	tx, err := storage.db.Beginx()
	err = tx.Get(&isAuthor, "SELECT true FROM UserProfile WHERE email = $1", senderEmail)
	if err != nil {
		return err
	}

	if !isAuthor {
		return errors.New("editing isn't allowed")
	}

	dtoToDB := map[string]string{
		"Text":             "text_content",
		"ShouldShowAuthor": "show_author",
		"ChangeDate":       "change_date",
	}
	dto.ChangeDate = utils.CurrentTimeString()
	err = repository.UpdateTable(storage.db, "Post", "id", *dto.PostID, dtoToDB, dto)
	if err != nil {
		return err
	}

	return nil
}

func (storage *PostStorage) DeletePost(senderEmail string, dto *dto.PostDelete) error {
	fmt.Println(dto)
	result, err := storage.db.Exec("UPDATE Post SET is_deleted = true WHERE author_id = (SELECT id FROM UserProfile WHERE email = $1) AND id = $2",
		senderEmail,
		dto.PostID)
	if err != nil {
		return err
	}

	deletedCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if deletedCount == 0 {
		return errors.New("resource doesn't exists or delete isn't allowed")
	}

	return nil
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
