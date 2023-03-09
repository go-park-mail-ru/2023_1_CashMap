package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
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
	if stored.Password != user.Password {
		return nil, apperror.IncorrectCredentials
	}

	return stored, nil
}

func (us *UserService) SignUp(user *entities.User) (*entities.User, error) {
	return us.repo.CreateUser(user)
}
