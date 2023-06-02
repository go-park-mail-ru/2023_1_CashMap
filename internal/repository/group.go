package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

//go:generate mockgen --destination=mocks/group.go depeche/internal/repository Group

type Group interface {
	GetGroupByLink(link string) (*entities.Group, error)
	GetUserGroupsByLink(link string, limit int, offset int) ([]*entities.Group, error)
	GetUserGroupsByEmail(email string, limit int, offset int) ([]*entities.Group, error)
	GetPopularGroups(email string, limit int, offset int) ([]*entities.Group, error)
	GetManagedGroups(email string, limit int, offset int) ([]*entities.Group, error)

	CreateGroup(ownerEmail string, group *dto.Group) (*entities.Group, error)
	UpdateGroup(link string, group *dto.UpdateGroup) error
	UpdateGroupAvatar(url, link string) error
	DeleteGroup(link string) error

	GetSubscribers(groupLink string, limit int, offset int) ([]*entities.User, error)
	Subscribe(email, groupLink string) error
	Unsubscribe(email, groupLink string) error
	AcceptRequest(userLink, groupLink string) error
	AcceptAllRequests(groupLink string) error
	DeclineRequest(userLink, groupLink string) error
	GetPendingRequests(groupLink string, limit, offset int) ([]*entities.User, error)

	IsOwner(userEmail, groupLink string) (bool, error)
	CheckSub(email, groupLink string) (bool, error)
	CheckAdmin(email, groupLink string) (bool, error)

	UpdateAvgGroupAvatarColor(color, link string) error

	// TODO добавить права доступа к группе
	// Grants(userEmail, groupLink string)
	// AddManager(manager *dto.AddManager) error
	// RemoveManager(link string) error
}
