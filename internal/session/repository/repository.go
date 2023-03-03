package repository

type Repository interface {
	CreateSession(token string, session *Session) error
	GetSession(token string) (*Session, error)
	DeleteSession(token string) error
}
