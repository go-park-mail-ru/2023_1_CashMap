package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type User interface {
	SignIn(user *dto.SignIn) (*entities.User, error)
	SignUp(user *dto.SignUp) (*entities.User, error)

	GetProfileByEmail(email string) (*dto.Profile, error)
	GetProfileByLink(link string) (*dto.Profile, error)

	Subscribe(subEmail, followLink string) error
	Unsubscribe(subEmail, followLink string) error
	Reject(rejectEmail, followerLink string) error

	GetFriendsByEmail(email string, limit, offset int) ([]*entities.User, error)
	GetSubscribesByEmail(email string, limit, offset int) ([]*entities.User, error)
	GetSubscribersByEmail(email string, limit, offset int) ([]*entities.User, error)

	GetFriendsByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error)
	GetSubscribesByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error)
	GetSubscribersByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error)
}
