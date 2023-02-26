package usecase

import "depeche/internal/entities"

type User interface {
	SignIn(user *entities.User) (string, error)
	SignUp(user *entities.User) (*entities.User, error)
	LogOut(token string) error
	CheckSession(token string) (bool, error)
}
