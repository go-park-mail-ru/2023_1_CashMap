package service

import (
	"depeche/authorization_ms/authEntities"
	"depeche/authorization_ms/repository"
	"depeche/pkg/apperror"
	"github.com/google/uuid"
	"time"
)

type Auth interface {
	Authenticate(email string) (string, error)
	LogOut(token string) error
	CheckSession(token string) (*authEntities.Session, error)
}

type AuthService struct {
	sessionRepo repository.SessionRepository
}

func NewAuthService(repo repository.SessionRepository) *AuthService {
	return &AuthService{
		sessionRepo: repo,
	}
}

func (a *AuthService) Authenticate(email string) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", apperror.InternalServerError
	}
	token := id.String()

	err = a.sessionRepo.CreateSession(token, &authEntities.Session{
		Email:     email,
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

func (a *AuthService) CheckSession(token string) (*authEntities.Session, error) {
	stored, err := a.sessionRepo.GetSession(token)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
