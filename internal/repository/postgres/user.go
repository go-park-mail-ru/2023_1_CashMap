package postgres

import (
	"depeche/internal/entities"
	"depeche/internal/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewPostgresUserRepo(db *sqlx.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
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
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
