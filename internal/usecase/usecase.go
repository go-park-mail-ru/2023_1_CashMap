package usecase

import (
	"depeche/internal/entities"
	"depeche/internal/session"
)

type User interface {
	SignIn(user *entities.User) (string, error)
	SignUp(user *entities.User) (*entities.User, error)
	LogOut(token string) error
	CheckSession(token string) (*session.Session, error)
}

type Feed interface {
	CollectPosts(user *entities.User, lastPostID uint, batchSize int, postsNumber int) ([]entities.Post, error)
}
