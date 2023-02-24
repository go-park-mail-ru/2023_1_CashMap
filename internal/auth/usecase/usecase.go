package usecase

import (
	"crypto/sha1"
	"depeche/internal/auth/entities"
	"depeche/internal/auth/repository"
	"errors"
)

type UseCase interface {
	AuthenticateUser(auth *entities.UserAuth) (string, error)
}

type AuthUseCase struct {
	Repository repository.Repository
}

func (useCase *AuthUseCase) AuthenticateUser(auth *entities.UserAuth) (string, error) {
	hasher := sha1.New()

	hasher.Write([]byte(auth.Password))
	passwordHash := string(hasher.Sum(nil))

	hashInStorage, err := useCase.Repository.GetPasswordHash(auth.Login)
	if err != nil {
		return "", err
	}

	if hashInStorage == passwordHash {
		return "Your hash", nil
	}

	return "", errors.New("invalid login or password")
}
