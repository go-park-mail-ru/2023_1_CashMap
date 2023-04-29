package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Group interface {
	GetGroupByLink(link string) (*entities.Group, error)
	GetUserGroupsByLink(link string, limit int, offset int) ([]*entities.Group, error)
	GetUSerGroupsByEmail(email string, limit int, offset int) ([]*entities.Group, error)
	GetPopularGroups(email string, limit int, offset int) ([]*entities.Group, error)

	CreateGroup(ownerEmail string, group *dto.Group) (*entities.Group, error)
	UpdateGroup(link string, group *dto.UpdateGroup) error
	DeleteGroup(link string) error
	Subscribe(email, groupLink string) error
	Unsubscribe(email, groupLink string) error
	AddManager(manager *dto.AddManager) error
	RemoveManager(link string) error
}
