package repository

import (
	session2 "depeche/authorization_ms"
)

type SessionRepository interface {
	CreateSession(token string, session *session2.Session) error
	GetSession(token string) (*session2.Session, error)
	DeleteSession(token string) error
}

type CSRFRepository interface {
	SaveCSRFToken(csrf *session2.CSRF, expirationTime int) error
	CheckCSRFToken(csrf *session2.CSRF) (bool, error)
	DeleteCSRFToken(csrf *session2.CSRF) error
}
