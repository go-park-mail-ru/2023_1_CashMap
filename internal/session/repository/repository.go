package repository

import (
	"depeche/internal/session"
)

type Repository interface {
	CreateSession(token string, session *session.Session) error
	GetSession(token string) (*session.Session, error)
	DeleteSession(token string) error
}
