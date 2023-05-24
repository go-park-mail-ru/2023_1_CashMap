package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type FeedRepository interface {
	GetFeedPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
	// GetFriendsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
	// GetGroupsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)

}
