package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"

	"github.com/fatih/structs"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewPostgresUserRepo(db *sqlx.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) GetUser(query string, args ...interface{}) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(query, args)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserById(id uint) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserByLink, id)
	err := row.StructScan(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByLink(link string) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserByLink, link)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserByEmail, email)
	err := row.StructScan(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.UserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetFriends(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(FriendsById, user.ID, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetSubscribes(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(SubscribesById, user.ID, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetSubscribers(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(SubscribesById, user.ID, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) Subscribe(subEmail, targetLink, requestTime string) (bool, error) {

	_, err := ur.DB.Queryx(Subscribe, subEmail, targetLink, requestTime)
	if err != nil {
		return false, err
	}
	return false, nil
}

func (ur *UserRepository) Unsubscribe(userEmail, targetLink string) (bool, error) {
	_, err := ur.DB.Queryx(Unsubscribe, userEmail, targetLink)
	if err != nil {
		return false, err
	}
	return false, nil
}

func (ur *UserRepository) RejectFriendRequest(userEmail, targetLink string) error {
	_, err := ur.DB.Queryx(Unsubscribe, userEmail, targetLink)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {

	err := ur.DB.QueryRowx(CreateUser, user.Email, user.Password).Err()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(email string, user *dto.EditProfile) (*entities.User, error) {
	query := "update userprofile set "
	var fields []interface{}
	for name, el := range structs.Map(user) {
		field, ok := el.(*string)
		if !ok {
			continue
		}
		if field != nil {
			fields = append(fields, *field)
			query += fmt.Sprintf("%s = $%d, ", mapNames[name], len(fields))
		}
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" where email = $%d", len(fields)+1)
	fields = append(fields, email)
	_, err := ur.DB.Queryx(query, fields...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

var mapNames = map[string]string{
	"Email":     "email",
	"Password":  "password",
	"FirstName": "first_name",
	"LastName":  "last_name",
	"Link":      "link",
	"Sex":       "sex",
	"Status":    "status",
	"Bio":       "bio",
	"Birthday":  "birthday",
}

func (ur *UserRepository) DeleteUser(email string, user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) GetPendingFriendRequests(user *entities.User, limit, offset int) ([]*dto.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) IsFriend(user, target *entities.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) IsSubscriber(user, target *entities.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) HasPendingRequest(user, target *entities.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) GetAllFriends(user *entities.User) ([]*dto.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) GetAllSubscribes(user *entities.User) ([]*dto.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) GetAllSubscribers(user *entities.User) ([]*dto.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) GetAllPendingFriendRequests(user *entities.User) ([]*dto.Profile, error) {
	//TODO implement me
	panic("implement me")
}
