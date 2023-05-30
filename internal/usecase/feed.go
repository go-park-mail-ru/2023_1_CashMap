package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

//go:generate mockgen --destination=mocks/feed.go depeche/internal/usecase Feed

type Feed interface {
	CollectPosts(email string, dto *dto.FeedDTO) ([]*entities.Post, error)
}
