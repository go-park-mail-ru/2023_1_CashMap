package local_storage

import (
	session "depeche/internal/session/repository"
	"errors"
)

type SessionStorage struct {
	session map[string]*session.Session
}

func NewMemorySessionStorage() *SessionStorage {
	return &SessionStorage{
		session: map[string]*session.Session{},
	}
}

func (s *SessionStorage) CreateSession(token string, session *session.Session) error {
	s.session[token] = session
	return nil
}

func (s *SessionStorage) GetSession(token string) (*session.Session, error) {
	stored := s.session[token]
	if stored == nil {
		return nil, errors.New("session not found")
	}
	return stored, nil
}

func (s *SessionStorage) DeleteSession(token string) error {
	stored := s.session[token]
	if stored == nil {
		return errors.New("session not found")
	}
	delete(s.session, token)
	return nil
}
