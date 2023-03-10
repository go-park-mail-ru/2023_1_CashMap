package local_storage

import (
	"depeche/internal/entities"
	"depeche/pkg/apperror"
	"sync"
)

type UserStorage struct {
	mx   sync.RWMutex
	user map[string]*entities.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		user: mockUsers,
	}
}

func (lc *UserStorage) GetUserById(id uint) (*entities.User, error) {
	lc.mx.Lock()
	defer lc.mx.Unlock()
	return nil, nil
}

func (lc *UserStorage) GetUserByEmail(email string) (*entities.User, error) {
	lc.mx.Lock()
	defer lc.mx.Unlock()
	user := lc.user[email]
	if user == nil {
		return nil, apperror.UserNotFound
	}
	return user, nil
}

func (lc *UserStorage) GetUserFriends(user *entities.User) ([]*entities.User, error) {
	lc.mx.Lock()
	defer lc.mx.Unlock()
	return nil, nil
}

func (lc *UserStorage) CreateUser(user *entities.User) (*entities.User, error) {
	lc.mx.RLock()
	defer lc.mx.RUnlock()
	if lc.user[user.Email] != nil {
		return nil, apperror.UserAlreadyExists
	}
	lc.user[user.Email] = user
	return user, nil
}

var mockUsers = map[string]*entities.User{
	"user1@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "Qwerty123!",
		FirstName: "Vladimir",
		LastName:  "Mayakovsky",
	},
	"user2@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Sergei",
		LastName:  "Esenin",
	},
	"user3@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Fedor",
		LastName:  "Tutchev",
	},
	"user4@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Michail",
		LastName:  "Lermontov",
	},
	"user5@mail.ru": {
		Email:     "user1@mail.ru",
		Password:  "some_hash",
		FirstName: "Alexandr",
		LastName:  "Pushkin",
	},
}
