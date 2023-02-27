package usecase

import (
	"depeche/internal/entities"
	"depeche/internal/session"
)

type User interface {
	SignIn(user *entities.User) (*entities.User, error)
	SignUp(user *entities.User) (*entities.User, error)
}

type Auth interface {
	Authenticate(user *entities.User) (string, error)
	LogOut(token string) error
	CheckSession(token string) (*session.Session, error)
}
