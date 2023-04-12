package redis

import (
	"depeche/internal/session"
	"depeche/internal/session/repository"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/go-redis/redis"
)

type Storage struct {
	Client         *redis.Client
	ExpirationTime int
}

func (s *Storage) CreateSession(token string, session *session.Session) error {
	_, err := s.Client.Do("SET", token, session.Email, "EX", s.ExpirationTime).Result()
	fmt.Println("NEW SESSION: ", session.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetSession(token string) (*session.Session, error) {
	email, err := s.Client.Get(token).Result()
	if err != nil {
		return nil, apperror.NoAuth
	}
	return &session.Session{
		Email: email,
	}, nil
}

func (s *Storage) DeleteSession(token string) error {
	_, err := s.Client.Do("DEL", token).Result()
	if err != nil {
		return err
	}
	return nil
}

func NewRedisStorage(client *redis.Client) repository.SessionRepository {
	return &Storage{
		Client:         client,
		ExpirationTime: 84600,
	}
}
