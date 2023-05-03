package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
)

type Group struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) usecase.Group {
	return &Group{
		repo: repo,
	}
}

func (g *Group) GetGroup(link string) (*entities.Group, error) {
	group, err := g.repo.GetGroupByLink(link)
	if err != nil {
		return nil, err
	}
	if !group.HideOwner {
		owner := entities.GroupManagement{
			Link: group.OwnerLink,
			Role: "owner",
		}
		group.Management = append(group.Management, owner)
	}
	return group, nil
}

func (g *Group) GetUserGroupsByLink(link string, limit int, offset int) ([]*entities.Group, error) {
	groups, err := g.repo.GetUserGroupsByLink(link, limit, offset)
	if err != nil {
		return nil, err
	}
	groups = addGroupManagement(groups)
	return groups, nil
}

func (g *Group) GetUserGroupsByEmail(email string, limit int, offset int) ([]*entities.Group, error) {
	groups, err := g.repo.GetUserGroupsByEmail(email, limit, offset)
	if err != nil {
		return nil, err
	}
	groups = addGroupManagement(groups)
	return groups, nil
}

func (g *Group) GetPopularGroups(email string, limit int, offset int) ([]*entities.Group, error) {
	groups, err := g.repo.GetPopularGroups(email, limit, offset)
	if err != nil {
		return nil, err
	}
	groups = addGroupManagement(groups)
	return groups, nil
}

func (g *Group) GetManagedGroups(email string, limit int, offset int) ([]*entities.Group, error) {
	groups, err := g.repo.GetManagedGroups(email, limit, offset)
	if err != nil {
		return nil, err
	}
	groups = addGroupManagement(groups)
	return groups, nil
}

func (g *Group) CreateGroup(ownerEmail string, group *dto.Group) error {
	_, err := g.repo.CreateGroup(ownerEmail, group)
	if err != nil {
		return err
	}
	return nil
}

func (g *Group) UpdateGroup(link string, ownerEmail string, group *dto.UpdateGroup) error {
	isOwner, err := g.repo.IsOwner(ownerEmail, link)
	if err != nil {
		return err
	}
	if !isOwner {
		return apperror.NewServerError(apperror.Forbidden, nil)
	}
	if group.Avatar != nil {
		err = g.repo.UpdateGroupAvatar(*group.Avatar, link)
		if err != nil {
			return err
		}
		group.Avatar = nil
	}

	if group.Link != nil {
		stored, err := g.repo.GetGroupByLink(*group.Link)
		if err != nil {
			return err
		}
		if stored != nil {
			return apperror.NewServerError(apperror.GroupAlreadyExists, nil)
		}
	}

	err = g.repo.UpdateGroup(link, group)
	if err != nil {
		return err
	}
	return nil
}

func (g *Group) DeleteGroup(ownerEmail string, link string) error {
	isOwner, err := g.repo.IsOwner(ownerEmail, link)
	if err != nil {
		return err
	}
	if !isOwner {
		return apperror.NewServerError(apperror.Forbidden, nil)
	}
	err = g.repo.DeleteGroup(link)
	if err != nil {
		return err
	}
	return nil
}

func (g *Group) GetSubscribers(groupLink string, limit int, offset int) ([]*entities.User, error) {
	users, err := g.repo.GetSubscribers(groupLink, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (g *Group) Subscribe(email, groupLink string) error {
	return g.repo.Subscribe(email, groupLink)
}

func (g *Group) Unsubscribe(email, groupLink string) error {
	return g.repo.Unsubscribe(email, groupLink)
}

func (g *Group) AcceptRequest(managerEmail, userLink, groupLink string) error {
	isOwner, err := g.repo.IsOwner(managerEmail, groupLink)
	if err != nil {
		return err
	}
	if !isOwner {
		return apperror.NewServerError(apperror.Forbidden, nil)
	}

	return g.repo.AcceptRequest(userLink, groupLink)
}

func (g *Group) AcceptAllRequests(managerEmail, groupLink string) error {
	isOwner, err := g.repo.IsOwner(managerEmail, groupLink)
	if err != nil {
		return err
	}
	if !isOwner {
		return apperror.NewServerError(apperror.Forbidden, nil)
	}

	return g.repo.AcceptAllRequests(groupLink)
}

func (g *Group) DeclineRequest(managerEmail, userLink, groupLink string) error {
	isOwner, err := g.repo.IsOwner(managerEmail, groupLink)
	if err != nil {
		return err
	}
	if !isOwner {
		return apperror.NewServerError(apperror.Forbidden, nil)
	}

	return g.repo.DeclineRequest(userLink, groupLink)
}

func (g *Group) GetPendingRequests(managerEmail, groupLink string, limit, offset int) ([]*entities.User, error) {
	isOwner, err := g.repo.IsOwner(managerEmail, groupLink)
	if err != nil {
		return nil, err
	}
	if !isOwner {
		return nil, apperror.NewServerError(apperror.Forbidden, nil)
	}

	return g.repo.GetPendingRequests(groupLink, limit, offset)
}

func (g *Group) CheckSub(email, groupLink string) (bool, error) {
	return g.repo.CheckSub(email, groupLink)
}

func (g *Group) CheckAdmin(email, groupLink string) (bool, error) {
	return g.repo.CheckAdmin(email, groupLink)
}

func addGroupManagement(groups []*entities.Group) []*entities.Group {
	for _, group := range groups {
		if !group.HideOwner {
			owner := entities.GroupManagement{
				Link: group.OwnerLink,
				Role: "owner",
			}
			group.Management = append(group.Management, owner)
		}
	}
	return groups
}
