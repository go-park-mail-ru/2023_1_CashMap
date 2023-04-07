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

	LikePost()
	CancelLike()

	UpdatePost(email string, dto *dto.PostUpdate) error

	Repost()

	//AddComment() - в comment service
	//RemoveComment()
	// TODO: сделать центр уведомлений для комментариев и сообщений входящий (да и для вообще всего)

}
