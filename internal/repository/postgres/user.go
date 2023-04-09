package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"

	"github.com/fatih/structs"
)

type UserRepository struct {
	DB *sqlx.DB
}

// TODO specify errors (replace 500 with smth) and find a way to wrap an sql error
// TODO remove println

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
		if err == sql.ErrNoRows {
			return nil, apperror.UserNotFound
		}
		return nil, apperror.InternalServerError
	}
	return user, nil
}

func (ur *UserRepository) GetUserById(id uint) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserById, id)
	err := row.StructScan(user)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, apperror.UserNotFound
		}
		return nil, apperror.InternalServerError
	}
	return user, nil
}

func (ur *UserRepository) GetUserByLink(link string) (*entities.User, error) {
	user := &entities.User{}
	row := ur.DB.QueryRowx(UserByLink, link)
	err := row.StructScan(user)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, apperror.UserNotFound
		}
		return nil, apperror.InternalServerError
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
		return nil, apperror.InternalServerError
	}
	return user, nil
}

func (ur *UserRepository) GetFriends(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(FriendsById, user.ID, limit, offset)
	if err != nil {
		return nil, apperror.InternalServerError
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.InternalServerError
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetSubscribes(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(SubscribesById, user.ID, limit, offset)
	if err != nil {
		return nil, apperror.InternalServerError
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.InternalServerError
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetSubscribers(user *entities.User, limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(SubscribersById, user.ID, limit, offset)
	if err != nil {
		return nil, apperror.InternalServerError
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.InternalServerError
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetUsers(limit, offset int) ([]*entities.User, error) {
	var users []*entities.User
	rows, err := ur.DB.Queryx(AllUsers, limit, offset)
	if err != nil {
		return nil, apperror.InternalServerError
	}
	for rows.Next() {
		var user = &entities.User{}
		if err := rows.StructScan(user); err != nil {
			return nil, apperror.InternalServerError
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) Subscribe(subEmail, targetLink, requestTime string) (bool, error) {
	_, err := ur.DB.Exec(Subscribe, subEmail, targetLink, requestTime)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case pgerrcode.UniqueViolation:
				return false, apperror.RepeatedSubscribe
			default:
				return false, apperror.InternalServerError
			}
		}
		return false, apperror.InternalServerError
	}
	return false, nil
}

func (ur *UserRepository) Unsubscribe(userEmail, targetLink string) (bool, error) {
	_, err := ur.DB.Queryx(Unsubscribe, userEmail, targetLink)
	if err != nil {
		fmt.Println(err)
		// TODO check subscribe conflict (repeated unsubscribe)
		return false, apperror.InternalServerError
	}
	return false, nil
}

func (ur *UserRepository) RejectFriendRequest(userEmail, targetLink string) error {
	_, err := ur.DB.Queryx(Unsubscribe, userEmail, targetLink)
	if err != nil {
		fmt.Println(err)
		// TODO check
		return apperror.InternalServerError

	}
	return nil
}

func (ur *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {

	tx, err := ur.DB.Beginx()
	if err != nil {
		fmt.Println(err)
		// TODO check
		return nil, apperror.InternalServerError
	}

	var id uint
	err = tx.QueryRowx(CreateUser, user.Email, user.Password, user.FirstName, user.LastName, utils.CurrentTimeString()).Scan(&id)
	if err != nil {
		fmt.Println(err)
		// TODO check
		return nil, apperror.InternalServerError

	}

	_, err = tx.Exec(UpdateUserLink, fmt.Sprintf("id%d", id), id)
	if err != nil {
		fmt.Println(err)
		// TODO check
		return nil, apperror.InternalServerError

	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		// TODO check
		return nil, apperror.InternalServerError

	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(email string, user *dto.EditProfile) (*entities.User, error) {
	query := "update userprofile set "
	// TODO: поправиить sql запрос
	var fields []interface{}
	for name, el := range structs.Map(user) {
		field, ok := el.(*string)
		if !ok {
			continue
		}
		if dbName, exists := mapNames[name]; field != nil && exists {
			fields = append(fields, *field)
			query += fmt.Sprintf("%s = $%d, ", dbName, len(fields))
		}
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" where email = $%d", len(fields)+1)
	fields = append(fields, email)
	_, err := ur.DB.Queryx(query, fields...)
	if err != nil {
		fmt.Println(err)
		return nil, apperror.InternalServerError
	}
	return nil, nil
}

var mapNames = map[string]string{
	"Email":       "email",
	"NewPassword": "password",
	"FirstName":   "first_name",
	"LastName":    "last_name",
	"Link":        "link",
	"Sex":         "sex",
	"Status":      "status",
	"Bio":         "bio",
	"Birthday":    "birthday",
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
