package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/session"
	"errors"
	"github.com/google/uuid"
	"time"
)

type UserService struct {
	repo        repository.UserRepository
	sessionRepo session.Repository
}

func NewUserService(repo repository.UserRepository, sRepo session.Repository) *UserService {
	return &UserService{
		repo:        repo,
		sessionRepo: sRepo,
	}
}

func (us *UserService) SignIn(user *entities.User) (string, error) {
	stored, err := us.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if stored.Password != user.Password {
		return "", errors.New("invalid login or password")
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	token := id.String()

	err = us.sessionRepo.CreateSession(token, &session.Session{
		Email:     user.Email,
		ExpiresAt: time.Now().Add(time.Second * 86400),
	})
	if err != nil {
		return "", err
	}
	return token, err
}

func (us *UserService) SignUp(user *entities.User) (*entities.User, error) {
	return us.repo.CreateUser(user)
}

func (us *UserService) LogOut(token string) error {
	err := us.sessionRepo.DeleteSession(token)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) CheckSession(token string) (*session.Session, error) {
	stored, err := us.sessionRepo.GetSession(token)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
