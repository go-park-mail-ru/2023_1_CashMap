package response

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type GetPostsByUserLinkResponse struct {
	Body GetPostsByUserLinkBody `json:"body"`
}

type GetPostsByUserLinkBody struct {
	Posts []*entities.Post `json:"posts"`
}

type GetPostsByCommunityLinkResponse struct {
	Body GetPostsByCommunityLinkBody `json:"body"`
}

type GetPostsByCommunityLinkBody struct {
	Posts []*entities.Post `json:"posts"`
}

type GetPostsByIdResponse struct {
	Body GetPostsByIdBody `json:"body"`
}

type GetPostsByIdBody struct {
	Posts []*entities.Post `json:"posts"`
}

type CreatePostRequest struct {
	Body *dto.PostCreate `json:"body"`
}

type CreatePostResponse struct {
	Body CreatePostBody `json:"body"`
}

type CreatePostBody struct {
	Posts []*entities.Post `json:"posts"`
}

type DeletePostRequest struct {
	Body *dto.PostDelete `json:"body"`
}

type LikePostRequest struct {
	Body *dto.LikeDTO `json:"body"`
}

type LikePostResponse struct {
	Body LikePostBody `json:"body"`
}

type LikePostBody struct {
	LikesAmount int `json:"likes_amount"`
}

type CancelPostLikeRequest struct {
	Body *dto.LikeDTO `json:"body"`
}
