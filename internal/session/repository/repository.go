package repository

import (
	"depeche/internal/session"
)

type SessionRepository interface {
	CreateSession(token string, session *session.Session) error
	GetSession(token string) (*session.Session, error)
	DeleteSession(token string) error
}

type CSRFRepository interface {
	SaveCSRFToken(csrf *session.CSRF, expirationTime int) error
	CheckCSRFToken(csrf *session.CSRF) (bool, error)
	DeleteCSRFToken(csrf *session.CSRF) error
}
