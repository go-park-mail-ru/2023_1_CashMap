package localmem

import (
	"errors"
)

type userLogin string

type passwordHash string

type AuthRepository struct {
	AuthStorage map[userLogin]passwordHash
}

func (repository *AuthRepository) GetPasswordHash(login string) (string, error) {
	if hash, exists := repository.AuthStorage[userLogin(login)]; exists {
		return string(hash), nil
	}

	return "", errors.New("user doesn't exists")
}
