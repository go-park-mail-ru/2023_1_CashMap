package service

import (
	"crypto/sha1"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"fmt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) usecase.User {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) SignIn(user *entities.User) (*entities.User, error) {
	stored, err := us.repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	passwordHash := hashPassword(user.Password)
	if stored.Password != passwordHash {
		return nil, apperror.IncorrectCredentials
	}

	return stored, nil
}

func (us *UserService) SignUp(user *entities.User) (*entities.User, error) {
	user.Password = hashPassword(user.Password)
	return us.repo.CreateUser(user)
}

func hashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	passwordHash := fmt.Sprintf("%x", hasher.Sum(nil)) // кодируем сырые байты в Base64
	return passwordHash
}
