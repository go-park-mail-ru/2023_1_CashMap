package redis

import (
	"depeche/internal/session"
	"depeche/internal/session/repository"
	"depeche/pkg/apperror"
	"github.com/go-redis/redis"
	"time"
)

var CSRF_TOKENS_HASH_MAP = "csrf"

var EXPIRATION_TIME = time.Second

type CSRFStorage struct {
	client *redis.Client
}

func NewCSRFStorage(client *redis.Client) repository.CSRFRepository {
	return &CSRFStorage{client}
}

// TODO: продумать нормальную эксприацию
func (storage *CSRFStorage) SaveCSRFToken(csrf *session.CSRF) error {
	_, err := storage.client.Do("HSET", CSRF_TOKENS_HASH_MAP, csrf.Token, csrf.Email).Result()
	if err != nil {
		return err
	}

	return nil
}

func (storage *CSRFStorage) CheckCSRFToken(csrf *session.CSRF) (bool, error) {
	email, err := storage.client.HGet(CSRF_TOKENS_HASH_MAP, csrf.Token).Result()
	if err != nil {
		return false, apperror.NoAuth
	}
	// TODO: ОТСЛЕДИТЬ, ЧТО ТАКОЙ ЗАПИСИ БОЛЬШЕ НЕТ И ВЕРНУТЬ ФАЛСЕ БЕЗ ОШИБКИ

	if email != csrf.Email {
		return false, apperror.NoAuth
	}

	return true, err
}

func (storage *CSRFStorage) DeleteCSRFToken(csrf *session.CSRF) error {
	email, err := storage.client.HGet(CSRF_TOKENS_HASH_MAP, csrf.Token).Result()
	if err != nil {
		return apperror.Forbidden
	}

	if email != csrf.Email {
		return apperror.Forbidden
	}

	_, err = storage.client.Do("HDEL", CSRF_TOKENS_HASH_MAP, csrf.Token).Result()
	if err != nil {
		return err
	}
	return nil
}
