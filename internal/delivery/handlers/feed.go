package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FeedHandler struct {
	service usecase.Feed
}

func NewFeedHandler(feedService usecase.Feed) *FeedHandler {
	return &FeedHandler{
		service: feedService,
	}
}

// GetFeed godoc
//
//	@Summary		Get feed part
//	@Description	Get user's new feed part by last post id and batch size.
//	@Tags			Feed
//	@Produce		json
//	@Param			batch_size		query		int		true	"Posts amount"
//	@Param			last_post_date	query		string	false	"Date and time of last post given. If not specified the newest posts will be sent"
//	@Success		200				{object}	doc.PostArray
//	@Failure		400				{object}	middleware.ErrorResponse
//	@Failure		401				{object}	middleware.ErrorResponse
//	@Failure		404				{object}	middleware.ErrorResponse
//	@Failure		500				{object}	middleware.ErrorResponse
//
//	@Router			/api/feed [get]
func (handler *FeedHandler) GetFeed(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	feedRequest := &dto.FeedDTO{}
	//err := ctx.ShouldBind(feedRequest)
	//if err != nil {
	//	_ = ctx.Error(apperror.BadRequest)
	//	return
	//}

	batchSize := ctx.Query("batch_size")
	bs, err := strconv.Atoi(batchSize)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
	}

	feedRequest.BatchSize = uint(bs)

	posts, err := handler.service.CollectPosts(email.(string), feedRequest)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"posts": posts,
		},
	})
}
