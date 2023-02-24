package repository

import "depeche/internal/auth/entities"

type Repository interface {
	CreateUser(user *entities.User) error
	GetPasswordHash(login string) (string, error)
	FindSession(sessionId string) (bool, error)
	CreateSession(session *entities.Session) error
}
