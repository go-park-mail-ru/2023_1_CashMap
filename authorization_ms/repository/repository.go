package repository

import (
	"depeche/authorization_ms/authEntities"
)

type SessionRepository interface {
	CreateSession(token string, session *authEntities.Session) error
	GetSession(token string) (*authEntities.Session, error)
	DeleteSession(token string) error
}

type CSRFRepository interface {
	SaveCSRFToken(csrf *authEntities.CSRF, expirationTime int) error
	CheckCSRFToken(csrf *authEntities.CSRF) (bool, error)
	DeleteCSRFToken(csrf *authEntities.CSRF) error
}
