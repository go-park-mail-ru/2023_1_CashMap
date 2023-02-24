package usecase

import (
	"crypto/sha1"
	"depeche/internal/auth/entities"
	"depeche/internal/auth/repository"
	"depeche/internal/auth/repository/localmem"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type AuthService interface {
	RegisterUser(user *entities.User) error
	AuthenticateUser(auth *entities.Credentials) (string, error)
	NewSessionToken(login string) (string, error)
	ValidateSession(sessionId string) (bool, error)
}

func NewAuthService() AuthService {
	return &SessionAuthService{
		Repository: localmem.NewRepository(),
	}
}

type SessionAuthService struct {
	Repository repository.Repository
}

func (useCase *SessionAuthService) ValidateSession(sessionId string) (bool, error) {
	ok, err := useCase.Repository.FindSession(sessionId)
	return ok, err
}

func (useCase *SessionAuthService) RegisterUser(user *entities.User) error {
	// TODO: валидация данных
	user.Password = hashPassword(user.Password)
	return useCase.Repository.CreateUser(user)
}

func (useCase *SessionAuthService) AuthenticateUser(auth *entities.Credentials) (string, error) {
	// TODO: запретить авторизацию, если сессия уже есть в базе
	passwordHash := hashPassword(auth.Password)

	hashFromStorage, err := useCase.Repository.GetPasswordHash(auth.Login)

	if err != nil {
		return "", err
	}

	if hashFromStorage == passwordHash {
		newToken, err := useCase.NewSessionToken(auth.Login)
		return newToken, err
	}

	return "", errors.New("invalid login or password")
}

func (useCase *SessionAuthService) NewSessionToken(login string) (string, error) {
	// TODO: метод New() может паниковать (судя из доки), обработать панику
	newId := uuid.New()
	err := useCase.Repository.CreateSession(&entities.Session{
		SessionId:      newId.String(),
		ExpirationTime: "2h",
		Login:          login,
	})

	if err != nil {
		return "", err
	}

	return newId.String(), nil
}

func hashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	passwordHash := fmt.Sprintf("%x", hasher.Sum(nil)) // кодируем сырые байты в Base64
	return passwordHash
}
