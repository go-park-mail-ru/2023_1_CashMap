package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type FeedRepository interface {
	GetFriendsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
	GetGroupsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error)
}
