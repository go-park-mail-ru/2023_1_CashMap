package _interface

import "depeche/internal/entities"

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
