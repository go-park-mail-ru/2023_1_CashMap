package redis

import (
	"depeche/authorization_ms/authEntities"
	"depeche/authorization_ms/repository"
	"depeche/pkg/apperror"
	"github.com/go-redis/redis"
)

type CSRFStorage struct {
	client *redis.Client
}

func NewCSRFStorage(client *redis.Client) repository.CSRFRepository {
	return &CSRFStorage{client}
}

func (storage *CSRFStorage) SaveCSRFToken(csrf *authEntities.CSRF, expirationTime int) error {
	_, err := storage.client.Do("SET", csrf.Token, csrf.Email, "EX", expirationTime).Result()
	if err != nil {
		return err
	}

	return nil
}

func (storage *CSRFStorage) CheckCSRFToken(csrf *authEntities.CSRF) (bool, error) {
	email, err := storage.client.Get(csrf.Token).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, apperror.InternalServerError
	}

	if email != csrf.Email {
		return false, apperror.NoAuth
	}

	return true, err
}

func (storage *CSRFStorage) DeleteCSRFToken(csrf *authEntities.CSRF) error {
	//  Проверка, что удаляем не чужой токен
	email, err := storage.client.Get(csrf.Token).Result()
	if err != nil {
		return apperror.Forbidden
	}

	if email != csrf.Email {
		return apperror.Forbidden
	}

	_, err = storage.client.Del(csrf.Token).Result()
	if err != nil {
		return err
	}
	return nil
}
