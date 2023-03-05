package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository/interface"
	"errors"
)

type User interface {
	SignIn(user *entities.User) (*entities.User, error)
	SignUp(user *entities.User) (*entities.User, error)
}

type UserService struct {
	repo _interface.UserRepository
}

func NewUserService(repo _interface.UserRepository) *UserService {
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
		return nil, errors.New("invalid login or password")
	}

	return stored, nil
}

func (us *UserService) SignUp(user *entities.User) (*entities.User, error) {
	return us.repo.CreateUser(user)
}
