package _interface

import (
	"depeche/internal/entities"
	"time"
)

// временное решение - потом заменим его на обращения к репам юзера и сообщества
type FeedRepository interface {
	GetFriendsPosts(user *entities.User, filterDateTime time.Time, postsNumber uint) ([]entities.Post, error)
	GetGroupsPosts(user *entities.User, filterDateTime time.Time, postsNumber uint) ([]entities.Post, error)
}
