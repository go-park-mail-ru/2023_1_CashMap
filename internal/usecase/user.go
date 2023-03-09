package usecase

import "depeche/internal/entities"

type User interface {
	SignIn(user *entities.User) (*entities.User, error)
	SignUp(user *entities.User) (*entities.User, error)
}
