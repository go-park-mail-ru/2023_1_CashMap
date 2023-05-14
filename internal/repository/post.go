package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type PostRepository interface {
	SelectPostById(postId uint, email string) (*entities.Post, error)
	SelectPostsByCommunityLink(info *dto.PostsGetByLink, email string) ([]*entities.Post, error)
	SelectPostsByUserLink(info *dto.PostsGetByLink, email string) ([]*entities.Post, error)

	GetPostSenderInfo(postID uint) (*entities.UserInfo, *entities.CommunityInfo, error)
	GetPostAttachments(postID uint) ([]string, error)

	CheckReadAccess(senderEmail string) (bool, error)
	CheckWriteAccess(senderEmail string, info *dto.PostCreate) (bool, error)

	CreatePost(senderEmail string, dto *dto.PostCreate) (uint, error)
	AddPostAttachments(postId uint, attachments []string) error
	UpdatePost(senderEmail string, dto *dto.PostUpdate) error
	DeletePost(senderEmail string, dto *dto.PostDelete) error

	SetLike(email string, postID uint) error
	CancelLike(email string, postID uint) error
	GetLikesAmount(email string, postID uint) (int, error)
}
