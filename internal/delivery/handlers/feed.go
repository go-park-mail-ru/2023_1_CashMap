package handlers

import (
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"github.com/gin-gonic/gin"
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

func (handler *FeedHandler) GetPosts(ctx *gin.Context) {
	feedRequest := struct {
		BatchSize    int       `json:"batch_size"`
		LastPostDate time.Time `json:"last_post_id"`
	}{}

	err := ctx.BindJSON(feedRequest)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		ctx.AbortWithError(400, err)
		return
	}

	user := &entities.User{
		Email: email.(string),
	}

	posts, err := handler.service.CollectPosts(user, feedRequest.LastPostDate, feedRequest.BatchSize)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(200, map[string][]entities.Post{
		"posts": posts,
	})

}
