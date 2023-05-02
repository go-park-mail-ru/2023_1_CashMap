package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	utildb "depeche/internal/repository/utils"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	"strings"
)

type GroupRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) repository.Group {
	return &GroupRepository{
		db: db,
	}
}

func (gr *GroupRepository) GetGroupByLink(link string) (*entities.Group, error) {
	group := &entities.Group{}
	err := gr.db.QueryRowx(GroupByLink, link).StructScan(group)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NewServerError(apperror.GroupNotFound, nil)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	return group, nil
}
func (gr *GroupRepository) GetUserGroupsByLink(link string, limit int, offset int) ([]*entities.Group, error) {
	var groups []*entities.Group
	rows, err := gr.db.Queryx(GroupsByUserlink, link, limit, offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return groups, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var group = &entities.Group{}
		if err := rows.StructScan(group); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (gr *GroupRepository) GetUserGroupsByEmail(email string, limit int, offset int) ([]*entities.Group, error) {
	var groups []*entities.Group
	rows, err := gr.db.Queryx(GroupsByUserEmail, email, limit, offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return groups, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var group = &entities.Group{}
		if err := rows.StructScan(group); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func (gr *GroupRepository) GetPopularGroups(email string, limit int, offset int) ([]*entities.Group, error) {
	var groups []*entities.Group
	// TODO добавить групы на которые не подписан
	rows, err := gr.db.Queryx(GetGroups, limit, offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return groups, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var group = &entities.Group{}
		if err := rows.StructScan(group); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (gr *GroupRepository) GetManagedGroups(email string, limit int, offset int) ([]*entities.Group, error) {
	var groups []*entities.Group
	rows, err := gr.db.Queryx(GetManaged, email, limit, offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return groups, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var group = &entities.Group{}
		if err := rows.StructScan(group); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (gr *GroupRepository) CreateGroup(ownerEmail string, group *dto.Group) (*entities.Group, error) {
	tx, err := gr.db.Beginx()
	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			err = fmt.Errorf("rollback error: %w", err)
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	var id uint
	err = tx.QueryRowx(CreateGroup, ownerEmail, group.Title, group.Info, group.Privacy, utils.CurrentTimeString(), group.HideOwner).Scan(&id)
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

var groupMapNames = map[string]string{
	"Title":     "title",
	"Info":      "info",
	"Privacy":   "privacy",
	"HideOwner": "hide_owner",
}

func (gr *GroupRepository) UpdateGroup(link string, group *dto.UpdateGroup) error {
	query := "update groups set "
	var fields []interface{}
	fieldsMap := structs.Map(group)
	for _, name := range structs.Names(group) {
		field, ok := fieldsMap[name].(*string)
		if !ok {
			continue
		}
		if dbName, exists := groupMapNames[name]; field != nil && exists {
			fields = append(fields, *field)
			query += fmt.Sprintf("%s = $%d, ", dbName, len(fields))
		}
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" where link = $%d", len(fields)+1)
	fields = append(fields, link)
	rows, err := gr.db.Queryx(query, fields...)
	defer utildb.CloseRows(rows)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	err = rows.Close()
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}

	return nil
}

func (gr *GroupRepository) UpdateGroupAvatar(url, link string) error {
	err := gr.db.QueryRowx(UpdateAvatar, url, link).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) DeleteGroup(link string) error {
	err := gr.db.QueryRowx(DeleteGroup, link).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) GetSubscribers(groupLink string, limit int, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := gr.db.Queryx(GroupSubscribers, groupLink, limit, offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (gr *GroupRepository) Subscribe(email, groupLink string) error {
	err := gr.db.QueryRowx(GroupSubscribe, email, groupLink).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) Unsubscribe(email, groupLink string) error {
	err := gr.db.QueryRowx(GroupUnsubscribe, email, groupLink).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) AcceptRequest(userLink, groupLink string) error {
	err := gr.db.QueryRowx(AcceptRequest, userLink, groupLink).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) AcceptAllRequests(groupLink string) error {
	err := gr.db.QueryRowx(AcceptAllRequests, groupLink).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) DeclineRequest(userLink, groupLink string) error {
	err := gr.db.QueryRowx(DeclineRequest, userLink, groupLink).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			return apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return nil
}

func (gr *GroupRepository) GetPendingRequests(groupLink string, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := gr.db.Queryx(PendingGroupRequests, groupLink, limit, offset)
	defer utildb.CloseRows(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (gr *GroupRepository) IsOwner(userEmail, groupLink string) (bool, error) {
	var isOwner bool
	err := gr.db.QueryRowx(IsOwner, userEmail, groupLink).Scan(&isOwner)
	if err != nil {
		if err != sql.ErrNoRows {
			return isOwner, apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return isOwner, nil
}

func (gr *GroupRepository) CheckSub(email, groupLink string) (bool, error) {
	var isSub bool
	err := gr.db.QueryRowx(CheckSub, email, groupLink).Scan(&isSub)
	if err != nil {
		if err != sql.ErrNoRows {
			return isSub, apperror.NewServerError(apperror.InternalServerError, err)
		}
	}
	return isSub, nil
}

//func (gr *GroupRepository) AddManager(manager *dto.AddManager) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (gr *GroupRepository) RemoveManager(link string) error {
//	//TODO implement me
//	panic("implement me")
//}
