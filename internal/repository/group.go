package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Group interface {
	CreateGroup(ownerEmail string, group *dto.CreateGroup) (*entities.Group, error)
	UpdateGroup(group *dto.UpdateGroup) error
	DeleteGroup(link string) error
	GroupsByEmail(email string) ([]*entities.Group, error)
	GroupsByLink(link string) ([]*entities.Group, error)
	Subscribe(email, groupLink string) error
	Unsubscribe(email, groupLink string) error
	AddManager(manager *dto.AddManager) error
	RemoveManager(link string) error
}
