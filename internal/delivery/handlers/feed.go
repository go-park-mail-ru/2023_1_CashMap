package handlers

import (
	"depeche/internal/delivery"
	"depeche/internal/entities"
	"depeche/internal/usecase/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedHandler struct {
	service service.Feed
}

func NewFeedHandler(feedService service.Feed) *FeedHandler {
	return &FeedHandler{
		service: feedService,
	}
}

// GetFeed godoc
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
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	feedRequest := struct {
		BatchSize    uint      `form:"batch_size"`
		LastPostDate time.Time `form:"last_post_id"`
	}{}

	err := ctx.ShouldBind(&feedRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := &entities.User{
		Email: email.(string),
	}

	posts, err := handler.service.CollectPosts(user, feedRequest.LastPostDate, feedRequest.BatchSize)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, delivery.Response{
		Body: map[string]interface{}{
			"post": posts,
		},
	})

}
