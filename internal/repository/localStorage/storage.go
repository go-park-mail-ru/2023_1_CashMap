package localStorage

import (
	"depeche/internal/entities"
	"errors"
)

type Storage struct {
	user map[string]*entities.User
}

func NewMemoryStorage() *Storage {
	return &Storage{
		user: map[string]*entities.User{},
	}
}

func (lc *Storage) GetUserById(id uint) (*entities.User, error) {
	return nil, nil
}
func (lc *Storage) GetUserByEmail(email string) (*entities.User, error) {
	user := lc.user[email]
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
func (lc *Storage) GetUserFriends(user *entities.User) ([]*entities.User, error) {
	return nil, nil
}
func (lc *Storage) CreateUser(user *entities.User) (*entities.User, error) {
	if lc.user[user.Email] != nil {
		return nil, errors.New("user already exists")
	}
	lc.user[user.Email] = user
	return user, nil
}
