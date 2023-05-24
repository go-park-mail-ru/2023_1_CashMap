package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Group interface {
	GetGroup(link string) (*entities.Group, error)
	GetUserGroupsByLink(link string, limit int, offset int) ([]*entities.Group, error)
	GetUserGroupsByEmail(email string, limit int, offset int) ([]*entities.Group, error)
	GetPopularGroups(email string, limit int, offset int) ([]*entities.Group, error)
	GetManagedGroups(email string, limit int, offset int) ([]*entities.Group, error)

	CreateGroup(ownerEmail string, group *dto.Group) error
	UpdateGroup(link string, ownerEmail string, group *dto.UpdateGroup) error
	DeleteGroup(ownerEmail string, link string) error

	GetSubscribers(groupLink string, limit int, offset int) ([]*entities.User, error)
	Subscribe(email, groupLink string) error
	Unsubscribe(email, groupLink string) error
	AcceptRequest(managerEmail, userLink, groupLink string) error
	AcceptAllRequests(managerEmail, groupLink string) error
	DeclineRequest(managerEmail, userLink, groupLink string) error
	GetPendingRequests(managerEmail, groupLink string, limit, offset int) ([]*entities.User, error)
	CheckSub(email, groupLink string) (bool, error)
	CheckAdmin(email, groupLink string) (bool, error)
}
