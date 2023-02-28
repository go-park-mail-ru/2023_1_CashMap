package delivery

import "github.com/gin-gonic/gin"

type FeedHandler interface {
	GetPosts(ctx *gin.Context)
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
}
