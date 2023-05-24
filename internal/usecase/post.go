package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type PostUsecase interface {
	GetPostById(email string, dto *dto.PostGetByID) (*entities.Post, error)
	GetPostsByCommunityLink(email string, dto *dto.PostsGetByLink) ([]*entities.Post, error)
	GetPostsByUserLink(email string, dto *dto.PostsGetByLink) ([]*entities.Post, error)

	CreatePost(email string, dto *dto.PostCreate) (*entities.Post, error)

	DeletePost(email string, dto *dto.PostDelete) error

	LikePost(email string, dto *dto.LikeDTO) (int, error)
	CancelLike(email string, dto *dto.LikeDTO) error

	UpdatePost(email string, dto *dto.PostUpdate) error

	Repost()

	AddPostData(post *entities.Post) error
}
