package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

//go:generate mockgen --destination=mocks/feed.go depeche/internal/repository FeedRepository

type FeedRepository interface {
	GetFriendsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
	GetGroupsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
	GetFeedPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
}
