package repository

import (
	"depeche/internal/entities"
	"time"
)

type FeedRepository interface {
	GetFriendsPosts(user *entities.User, filterDateTime time.Time, postsNumber int) ([]entities.Post, error)
	GetGroupsPosts(user *entities.User, filterDateTime time.Time, postsNumber int) ([]entities.Post, error)
}

type UserRepository interface {
	GetUserById(id uint) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserFriends(user *entities.User) ([]*entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
	/*
		UpdateUser(user *entities.User) error
		DeleteUser(id uint) error
	*/
}
