package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type CommentUsecase interface {
	GetCommentById(email string, id uint) (*entities.Comment, error)
	GetCommentsByPostId(email string, dto *dto.GetCommentsDTO) ([]*entities.Comment, bool, error)

	CreateComment(email string, dto *dto.CreateCommentDTO) (*entities.Comment, error)
	EditComment(email string, dto *dto.EditCommentDTO) error
	DeleteComment(email string, id uint) error
}
