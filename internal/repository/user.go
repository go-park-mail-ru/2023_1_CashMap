package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type UserRepository interface {
	GetUser(query string, args ...interface{}) (*entities.User, error)

	GetUserById(id uint) (*entities.User, error)
	GetUserByLink(link string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)

	GetFriends(user *entities.User, limit, offset int) ([]*entities.User, error)
	GetSubscribes(user *entities.User, limit, offset int) ([]*entities.User, error)
	GetSubscribers(user *entities.User, limit, offset int) ([]*entities.User, error)
	GetPendingFriendRequests(user *entities.User, limit, offset int) ([]*entities.User, error)
	GetUsers(email string, limit, offset int) ([]*entities.User, error)

	UpdateAvatar(email string, url string) error

	CheckLinkExists(link string) (bool, error)

	Subscribe(subEmail, targetLink, requestTime string) (bool, error)
	Unsubscribe(userEmail, targetLink string) (bool, error)
	RejectFriendRequest(userEmail, targetLink string) error

	//IsFriend(user, target *entities.User) (bool, error)
	//IsSubscriber(user, target *entities.User) (bool, error)
	//HasPendingRequest(user, target *entities.User) (bool, error)

	CreateUser(user *entities.User) (*entities.User, error)
	UpdateUser(email string, user *dto.EditProfile) (*entities.User, error)
	DeleteUser(email string) error

	SearchUserByName(email string, searchDTO *dto.GlobalSearchDTO) ([]*entities.UserInfo, error)
}
