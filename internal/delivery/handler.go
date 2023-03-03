package delivery

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	FeedHandler
	PostHandler
	UserHandler
}

type FeedHandler interface {
	GetFeed(ctx *gin.Context)
}

type PostHandler interface {
	GetPost(ctx *gin.Context)
	GetPostsBatch(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	EditPost(ctx *gin.Context)
}

type UserHandler interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	LogOut(ctx *gin.Context)
	CheckAuth(ctx *gin.Context)
}

type Response struct {
	Body map[string]interface{} `json:"body"`
}
