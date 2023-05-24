package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities"
	"depeche/internal/entities/response"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
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

	_response := response.GetPostsByUserLinkResponse{
		Body: response.GetPostsByUserLinkBody{
			Posts: posts,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
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

	_response := response.GetPostsByUserLinkResponse{
		Body: response.GetPostsByUserLinkBody{
			Posts: posts,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
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

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var post *entities.Post
	post, err = handler.PostUsecase.GetPostById(email, &postDTO)
	if err != nil {
		fmt.Println(err)
		_ = ctx.Error(apperror.InternalServerError)
		return
	}

	_response := response.GetPostsByUserLinkResponse{
		Body: response.GetPostsByUserLinkBody{
			Posts: []*entities.Post{post},
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

// CreatePost godoc
//
//	@Summary		Create new post
//	@Description	User can create new post
//	@Tags			Post
//	@Produce		json
//	@Param			request	formData	dto.PostCreate	true	"New post info"
//	@Success		200		{object}	doc.PostsResponse
//	@Failure		401
//	@Failure		500
//	@Router			/api/posts/create [post]
func (handler *PostHandler) CreatePost(ctx *gin.Context) {
	inputDTO := new(response.CreatePostRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	post, err := handler.PostUsecase.CreatePost(email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetPostsByUserLinkResponse{
		Body: response.GetPostsByUserLinkBody{
			Posts: []*entities.Post{post},
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
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
	inputDTO := new(response.DeletePostRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = handler.PostUsecase.DeletePost(email, inputDTO.Body)
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
//	@Param			request	formData	dto.PostUpdate	false	"Post to update data"
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

// LikePost godoc
//
//	@Summary		Set like on post
//	@Description	User can like posts if like hasn't already set
//	@Tags			Post
//	@Produce		json
//	@Param			request	body	dto.LikeDTO	 true	"Post data to like"
//	@Success		200		{object}	doc.LikePost
//	@Failure		400		{object} middleware.ErrorResponse
//	@Failure		401		{object} middleware.ErrorResponse
//	@Failure		409		{object} middleware.ErrorResponse
//	@Failure		500		{object} middleware.ErrorResponse
//	@Router			/api/posts/like/set [post]
func (handler *PostHandler) LikePost(ctx *gin.Context) {
	inputDTO := new(response.LikePostRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NewServerError(apperror.NoAuth, errors.New("failed to get email from context")))
		return
	}

	newLikesAmount, err := handler.PostUsecase.LikePost(email.(string), inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.LikePostResponse{
		Body: response.LikePostBody{
			LikesAmount: newLikesAmount,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

// CancelPostLike godoc
//
//	@Summary		Cancel post like
//	@Description	User can deny post like if like exists
//	@Tags			Post
//	@Produce		json
//	@Param			request	body	dto.LikeDTO	 true	"Post data to cancel like"
//	@Success		200
//	@Failure		400		{object} middleware.ErrorResponse
//	@Failure		401		{object} middleware.ErrorResponse
//	@Failure		409		{object} middleware.ErrorResponse
//	@Failure		500		{object} middleware.ErrorResponse
//	@Router			/api/posts/like/cancel [post]
func (handler *PostHandler) CancelPostLike(ctx *gin.Context) {
	inputDTO := new(response.CancelPostLikeRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		_ = ctx.Error(apperror.NewServerError(apperror.NoAuth, errors.New("failed to get email from context")))
		return
	}

	err := handler.PostUsecase.CancelLike(email.(string), inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
