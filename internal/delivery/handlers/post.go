package handlers

import (
	"depeche/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	usecase.PostUsecase
}

func NewPostHandler(postService usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		postService,
	}
}

func GetPost(ctx *gin.Context) {

}

func GetPostsBatch(ctx *gin.Context) {

}

func DeletePost(ctx *gin.Context) {

}

func EditPost(ctx *gin.Context) {

}
