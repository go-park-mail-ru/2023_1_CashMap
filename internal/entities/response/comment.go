package response

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

//go:generate easyjson --all comment.go

type GetCommentByIdResponse struct {
	Body GetCommentByIdBody `json:"body"`
}

type GetCommentByIdBody struct {
	Comment *entities.Comment `json:"comment"`
}

type GetCommentByPostIdResponse struct {
	Body GetCommentByPostIdBody `json:"body"`
}

type GetCommentByPostIdBody struct {
	Comments []*entities.Comment `json:"comments"`
	HasNext  bool                `json:"has_next"`
}

type CreateCommentRequest struct {
	Body *dto.CreateCommentDTO `json:"body"`
}

type CreateCommentResponse struct {
	Body CreateCommentBody `json:"body"`
}

type CreateCommentBody struct {
	Comment *entities.Comment `json:"comment"`
}

type EditCommentRequest struct {
	Body *dto.EditCommentDTO `json:"body"`
}
