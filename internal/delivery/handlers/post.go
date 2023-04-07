package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostHandler struct {
	usecase.PostUsecase
}

func NewPostHandler(postService usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		postService,
	}
}

// GetPostsByUserLink godoc
//
//	@Summary		Get post by user link
//	@Description	User can get user's posts in includes batches older than specified in "last_post_date"
//	@Tags			Post
//	@Param			owner_link		query	uint	true	"ID of the user on whose wall the post is located"
//	@Param			batch_size		query	uint	true	"Posts amount"
//	@Param			last_post_date	query	string	false	"Date and time of last post given. If not specified the newest posts will be sent"
//	@Success		200				{array}	doc.PostsResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/profile [get]
func (handler *PostHandler) GetPostsByUserLink(ctx *gin.Context) {
	post := dto.PostsGetByLink{}
	err := ctx.ShouldBind(&post)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	var posts []*entities.Post
	posts, err = handler.PostUsecase.GetPostsByUserLink(email.(string), &post)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"posts":       posts,
			"attachments": nil,
		},
	})
}

// GetPostsByCommunityLink godoc
//
//	@Summary		Get post by community link
//	@Description	User can get community's posts in includes batches older than specified in "last_post_date"
//	@Tags			Post
//	@Param			community_link	query	uint	true	"ID of the community on whose wall the post is located"
//	@Param			batch_size		query	uint	true	"Posts amount"
//	@Param			last_post_date	query	string	false	"Date and time of last post given. If not specified the newest posts will be sent"
//	@Success		200				{array}	doc.PostsResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/community [get]
func (handler *PostHandler) GetPostsByCommunityLink(ctx *gin.Context) {

	post := dto.PostsGetByLink{}
	err := ctx.ShouldBind(&post)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	if post.CommunityLink == nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	var posts []*entities.Post
	posts, err = handler.PostUsecase.GetPostsByCommunityLink(email.(string), &post)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"posts": posts,
		},
	})
}

// GetPostsById godoc
//
//	@Summary		Get post by id
//	@Description	User can get post by id, returned by server from CreatePost handler
//	@Tags			Post
//	@Param			post_id			query		uint	true	"Post ID"
//	@Param			batch_size		query		uint	true	"Posts amount"
//	@Param			last_post_date	query		string	false	"Date and time of last post given. If not specified the newest posts will be sent"
//	@Success		200				{object}	doc.PostResponse
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/id [get]
func (handler *PostHandler) GetPostsById(ctx *gin.Context) {
	postDTO := dto.PostGetByID{}

	err := ctx.ShouldBind(&postDTO)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	var post *entities.Post
	post, err = handler.PostUsecase.GetPostById(email.(string), &postDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"posts": []*entities.Post{post},
		},
	})
}

// CreatePost godoc
//
//	@Summary		Create new post
//	@Description	User can create new post
//	@Tags			Post
//	@Produce		json
//	@Param			request	formData	doc.PostCreateRequest	true	"New post info"
//	@Success		200		{object}	doc.PostsResponse
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/create [post]
func (handler *PostHandler) CreatePost(ctx *gin.Context) {
	var request = struct {
		dto.PostCreate `json:"body"`
	}{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		fmt.Println(err)
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	post, err := handler.PostUsecase.CreatePost(email.(string), &request.PostCreate)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"posts": []*entities.Post{post},
		},
	})
}

// DeletePost godoc
//
//	@Summary		Delete post by id
//	@Description	User can delete post
//	@Tags			Post
//	@Param			request	body	doc.PostDelete	false	"Post to delete info"
//	@Success		200
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/delete [delete]
func (handler *PostHandler) DeletePost(ctx *gin.Context) {
	var request = struct {
		dto.PostDelete `json:"body"`
	}{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}
	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	err = handler.PostUsecase.DeletePost(email.(string), &request.PostDelete)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

// EditPost godoc
//
//	@Summary		Edit post by id
//	@Description	User can edit post
//	@Tags			Post
//	@Param			request	formData	doc.PostUpdateRequest	false	"Post to update data"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/edit [patch]
func (handler *PostHandler) EditPost(ctx *gin.Context) {
	postUpdate := dto.PostUpdate{}
	err := ctx.ShouldBind(&postUpdate)
	if err != nil {
		_ = ctx.Error(apperror.BadRequest)
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		fmt.Println(err)
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	err = handler.UpdatePost(email.(string), &postUpdate)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
