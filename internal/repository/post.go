package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type PostRepository interface {
	SelectPostById(postId uint) (*entities.Post, error)
	SelectPostsByCommunityLink(info *dto.PostsGetByLink) ([]*entities.Post, error)
	SelectPostsByUserLink(info *dto.PostsGetByLink) ([]*entities.Post, error)

	CheckReadAccess(senderEmail string) (bool, error)
	CheckWriteAccess(senderEmail string, info *dto.PostCreate) (bool, error)

	CreatePost(senderEmail string, dto *dto.PostCreate) (uint, error)
	UpdatePost(senderEmail string, dto *dto.PostUpdate) error
	DeletePost(senderEmail string, dto *dto.PostDelete) error
}
