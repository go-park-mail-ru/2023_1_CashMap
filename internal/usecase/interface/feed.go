package _interface

import (
	"depeche/internal/entities"
	"time"
)

type Feed interface {
	CollectPosts(user *entities.User, lastPostDate time.Time, batchSize uint) ([]entities.Post, error)
}
