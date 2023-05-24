package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities/response"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	service usecase.CommentUsecase
}

func NewCommentHandler(commentService usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		service: commentService,
	}
}

func (gh *CommentHandler) GetCommentById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, fmt.Errorf("given id from url isn't valid uint: %s", ctx.Param("id"))))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	comment, err := gh.service.GetCommentById(email, uint(id))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetCommentByIdResponse{
		Body: response.GetCommentByIdBody{
			Comment: comment,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *CommentHandler) GetCommentByPostId(ctx *gin.Context) {
	postId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, fmt.Errorf("given post_id from url isn't valid uint: %s", ctx.Param("post_id"))))
		return
	}

	var request = struct {
		dto.GetCommentsDTO
	}{}

	err = ctx.ShouldBind(&request)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse GetCommentsDTO")))
		return
	}

	request.ID = uint(postId)

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	comments, hasNext, err := gh.service.GetCommentsByPostId(email, &request.GetCommentsDTO)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetCommentByPostIdResponse{
		Body: response.GetCommentByPostIdBody{
			Comments: comments,
			HasNext:  hasNext,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *CommentHandler) CreateComment(ctx *gin.Context) {
	inputDTO := new(response.CreateCommentRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	comment, err := gh.service.CreateComment(email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.CreateCommentResponse{
		Body: response.CreateCommentBody{
			Comment: comment,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (gh *CommentHandler) EditComment(ctx *gin.Context) {
	inputDTO := new(response.EditCommentRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	err = gh.service.EditComment(email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (gh *CommentHandler) DeleteComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, fmt.Errorf("given id from url isn't valid uint: %s", ctx.Param("id"))))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	err = gh.service.DeleteComment(email, uint(id))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
