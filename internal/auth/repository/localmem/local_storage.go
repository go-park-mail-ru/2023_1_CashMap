package localmem

import (
	"depeche/internal/auth/entities"
	"depeche/internal/auth/repository"
	"errors"
)

type userLogin string

type passwordHash string

type hash string

type AuthRepository struct {
	UsersStorage    []entities.User
	SessionsStorage []entities.Session
}

func NewRepository() repository.Repository {
	users := []entities.User{
		{
			entities.Credentials{
				Login:    "login",
				Password: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", // password
			},
			"Annoying",
			"Donkey",
			"men",
			"10.12.2001",
		},
		{
			entities.Credentials{
				Login:    "pavlen",
				Password: "b1b3773a05c0ed0176787a4f1574ff0075f7521e", // qwerty
			},
			"Pavel",
			"Repin",
			"men",
			"17.07.2003",
		},
		{
			entities.Credentials{
				Login:    "egor",
				Password: "b2751ceba6ac7c3dc2533b1112eb3fb36c107a00", // ne_boley
			},
			"firstName",
			"lastName",
			"men",
			"01.01.2001",
		},
	}

	return &AuthRepository{
		UsersStorage:    users,
		SessionsStorage: []entities.Session{},
	}
}

func (repository *AuthRepository) CreateSession(session *entities.Session) error {
	repository.SessionsStorage = append(repository.SessionsStorage, *session)
	return nil
}

func (repository *AuthRepository) CreateUser(user *entities.User) error {
	repository.UsersStorage = append(repository.UsersStorage, *user)
	return nil
}

func (repository *AuthRepository) GetPasswordHash(login string) (string, error) {
	for _, user := range repository.UsersStorage {
		if user.Login == login {
			return user.Password, nil
		}
	}

	return "", errors.New("user doesn't exists")
}

func (repository *AuthRepository) FindSession(sessionId string) (bool, error) {

	for _, session := range repository.SessionsStorage {

		if session.SessionId == sessionId {
			return true, nil
		}
	}

	return false, nil
}
