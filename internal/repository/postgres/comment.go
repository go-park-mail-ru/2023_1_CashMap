package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"github.com/jmoiron/sqlx"
)

type CommentStorage struct {
	db *sqlx.DB
}

func NewCommentStorage(db *sqlx.DB) repository.CommentRepository {
	return &CommentStorage{
		db: db,
	}
}

func (storage CommentStorage) GetCommentById(email string, id uint) (*entities.Comment, error) {
	var comment entities.Comment
	err := storage.db.Get(&comment, GetCommentByIdQuery, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	comment.SenderInfo = new(entities.CommentSenderInfo)
	err = storage.db.Get(comment.SenderInfo, GetCommentSenderInfoQuery, comment.ID)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	if comment.ReplyReceiver != nil {
		comment.ReplyReceiver = new(entities.CommentSenderInfo)
		err = storage.db.Get(comment.ReplyReceiver, GetReplyReceiverInfoQuery, comment.ID)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
	}

	return &comment, nil
}

func (storage CommentStorage) GetCommentsByPostId(email string, dto *dto.GetCommentsDTO) ([]*entities.Comment, bool, error) {
	var comments []*entities.Comment

	err := storage.db.Select(&comments, GetCommentsByPostIdQuery, dto.ID, dto.Count, dto.LastCommentDate, email)
	if err == sql.ErrNoRows {
		return comments, false, nil
	}
	if err != nil {
		return nil, false, apperror.NewServerError(apperror.InternalServerError, err)
	}

	for _, comment := range comments {
		comment.SenderInfo = new(entities.CommentSenderInfo)
		err = storage.db.Get(comment.SenderInfo, GetCommentSenderInfoQuery, comment.ID)
		if err != nil {
			return nil, false, apperror.NewServerError(apperror.InternalServerError, err)
		}

		if comment.ReplyReceiver != nil {
			comment.ReplyReceiver = new(entities.CommentSenderInfo)
			err = storage.db.Get(comment.ReplyReceiver, GetReplyReceiverInfoQuery, comment.ID)
			if err != nil {
				return nil, false, apperror.NewServerError(apperror.InternalServerError, err)
			}
		}
	}
	var hasNext bool
	if len(comments) != 0 {
		err = storage.db.Get(&hasNext, HasNextCommentsQuery, comments[len(comments)-1].CreationDate, comments[len(comments)-1].PostID)
		if err != nil {
			return nil, false, err
		}
	} else {
		hasNext = false
	}

	return comments, hasNext, nil
}

func (storage CommentStorage) CreateComment(email string, dto *dto.CreateCommentDTO) (uint, error) {
	var id uint
	err := storage.db.QueryRowx(CreateCommentQuery, dto.PostID, email, dto.ReplyTo, dto.Text, utils.CurrentTimeString()).Scan(&id)
	if err != nil {
		return 0, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return id, nil
}

func (storage CommentStorage) EditComment(email string, dto *dto.EditCommentDTO) error {
	_, err := storage.db.Exec(UpdateCommentQuery, dto.ID, dto.Text, utils.CurrentTimeString())
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil
}

func (storage CommentStorage) DeleteComment(email string, id uint) error {
	var isDeleted bool
	err := storage.db.Get(&isDeleted, IsCommentDeleted, id)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	if isDeleted {
		return apperror.NewServerError(apperror.Forbidden, errors.New("comment has already been deleted"))
	}

	_, err = storage.db.Exec(DeleteCommentQuery, id)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil
}

func (storage CommentStorage) IsCommentAuthor(id uint, email string) (bool, error) {
	var isAuthor bool
	err := storage.db.Get(&isAuthor, IsCommentAuthor, id, email)
	if err != nil {
		return false, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return isAuthor, nil
}
