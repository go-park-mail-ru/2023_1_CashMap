package service

import (
	"depeche/internal/entities"
	"depeche/internal/session"
	"depeche/internal/session/repository"
	"depeche/pkg/apperror"
	"github.com/google/uuid"
	"time"
)

type Auth interface {
	Authenticate(user *entities.User) (string, error)
	LogOut(token string) error
	CheckSession(token string) (*session.Session, error)
}

type AuthService struct {
	sessionRepo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{
		sessionRepo: repo,
	}
}

func (a *AuthService) Authenticate(user *entities.User) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", apperror.InternalServerError
	}
	token := id.String()

	err = a.sessionRepo.CreateSession(token, &session.Session{
		Email:     user.Email,
		ExpiresAt: time.Now().Add(time.Second * 86400),
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) LogOut(token string) error {
	err := a.sessionRepo.DeleteSession(token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) CheckSession(token string) (*session.Session, error) {
	stored, err := a.sessionRepo.GetSession(token)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
