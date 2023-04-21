package postgres

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GroupRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) repository.Group {
	return &GroupRepository{
		db: db,
	}
}

func (gr *GroupRepository) CreateGroup(ownerEmail string, group *dto.CreateGroup) (*entities.Group, error) {
	tx, err := gr.db.Beginx()
	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			err = fmt.Errorf("rollback error: %w", err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	var id uint
	err = tx.QueryRowx(CreateGroup, group.Title, group.Info, ownerEmail, group.HideOwner, utils.CurrentTimeString()).Scan(&id)
	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			err = fmt.Errorf("rollback error: %w", err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	_, err = tx.Exec(UpdateGroupLink, fmt.Sprintf("id%d", id), id)
	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			err = fmt.Errorf("rollback error: %w", err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	if group.Avatar != "" {
		_, err = tx.Exec(UpdateGroupAvatar, group.Avatar, id)
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				err = fmt.Errorf("rollback error: %w", err)
			}
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			err = fmt.Errorf("rollback error: %w", err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return nil, nil
}

func (gr *GroupRepository) UpdateGroup(group *dto.UpdateGroup) error {

	return nil
}

func (gr *GroupRepository) DeleteGroup(link string) error {
	//TODO implement me
	panic("implement me")
}

func (gr *GroupRepository) GroupsByEmail(email string) ([]*entities.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (gr *GroupRepository) GroupsByLink(link string) ([]*entities.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (gr *GroupRepository) Subscribe(email, groupLink string) error {
	//TODO implement me
	panic("implement me")
}

func (gr *GroupRepository) Unsubscribe(email, groupLink string) error {
	//TODO implement me
	panic("implement me")
}

func (gr *GroupRepository) AddManager(manager *dto.AddManager) error {
	//TODO implement me
	panic("implement me")
}

func (gr *GroupRepository) RemoveManager(link string) error {
	//TODO implement me
	panic("implement me")
}
