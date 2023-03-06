package repository

import (
	"depeche/internal/entities"
	"time"
)

type FeedRepository interface {
	GetFriendsPosts(user *entities.User, filterDateTime time.Time, postsNumber uint) ([]entities.Post, error)
	GetGroupsPosts(user *entities.User, filterDateTime time.Time, postsNumber uint) ([]entities.Post, error)
}
