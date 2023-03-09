package handlers

import (
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
//	@Description	Get users's new feed part by last post id and batch size.
//	@Tags			feed
//	@Produce		json
//	@Param			batch_size		query	int		true	"Posts amount"
//	@Param			last_post_date	query	string	false	"Date and time of last post given. If not specified the newest posts will be sent"
//	@Success		200				{array}	entities.Post
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			api/feed [get]
func (handler *FeedHandler) GetFeed(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	feedRequest := struct {
		BatchSize    uint      `form:"batch_size"`
		LastPostDate time.Time `form:"last_post_id"`
	}{}

	err := ctx.ShouldBind(&feedRequest)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	user := &entities.User{
		Email: email.(string),
	}

	posts, err := handler.service.CollectPosts(user, feedRequest.LastPostDate, feedRequest.BatchSize)
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
