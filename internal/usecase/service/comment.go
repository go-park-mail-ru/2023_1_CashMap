package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) usecase.CommentUsecase {
	return CommentService{
		repo: repository,
	}
}

func (c CommentService) GetCommentById(email string, id uint) (*entities.Comment, error) {
	comment, err := c.repo.GetCommentById(email, id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c CommentService) GetCommentsByPostId(email string, dto *dto.GetCommentsDTO) ([]*entities.Comment, bool, error) {
	if dto.LastCommentDate == nil {
		oldestDate := "0"
		dto.LastCommentDate = &oldestDate
	}

	comments, hasNext, err := c.repo.GetCommentsByPostId(email, dto)
	if err != nil {
		return nil, false, err
	}

	return comments, hasNext, nil
}

func (c CommentService) CreateComment(email string, dto *dto.CreateCommentDTO) (*entities.Comment, error) {
	if dto.PostID == nil {
		return nil, apperror.NewServerError(apperror.BadRequest, errors.New("postID is required field"))
	}
	commentID, err := c.repo.CreateComment(email, dto)
	if err != nil {
		return nil, err
	}

	return c.GetCommentById(email, commentID)
}

func (c CommentService) EditComment(email string, dto *dto.EditCommentDTO) error {
	if dto.ID == nil {
		return apperror.NewServerError(apperror.BadRequest, errors.New("id is required field"))
	}
	isAuthor, err := c.repo.IsCommentAuthor(*dto.ID, email)
	if err != nil {
		return err
	}

	if !isAuthor {
		return apperror.NewServerError(apperror.Forbidden, errors.New("editing foreign comments is not allowed"))
	}

	err = c.repo.EditComment(email, dto)
	if err != nil {
		return err
	}

	return nil
}

func (c CommentService) DeleteComment(email string, id uint) error {
	isAuthor, err := c.repo.IsCommentAuthor(id, email)
	if err != nil {
		return err
	}

	if !isAuthor {
		return apperror.NewServerError(apperror.Forbidden, errors.New("deleting foreign comments is not allowed"))
	}

	err = c.repo.DeleteComment(email, id)
	if err != nil {
		return err
	}

	return nil
}
