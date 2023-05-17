package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"comment": comment,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"comments": comments,
			"has_next": hasNext,
		},
	})
}

func (gh *CommentHandler) CreateComment(ctx *gin.Context) {
	var request = struct {
		*dto.CreateCommentDTO `json:"body"`
	}{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse CreateCommentDTO")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	comment, err := gh.service.CreateComment(email, request.CreateCommentDTO)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"comment": comment,
		},
	})
}

func (gh *CommentHandler) EditComment(ctx *gin.Context) {
	var request = struct {
		*dto.EditCommentDTO `json:"body"`
	}{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse EditCommentDTO")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
	}

	err = gh.service.EditComment(email, request.EditCommentDTO)
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
