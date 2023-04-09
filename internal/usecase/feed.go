package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Feed interface {
	CollectPosts(email string, dto *dto.FeedDTO) ([]*entities.Post, error)
}
