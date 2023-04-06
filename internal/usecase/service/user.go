package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) usecase.User {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) SignIn(user *dto.SignIn) (*entities.User, error) {
	stored, err := us.repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if stored.Password != utils.Hash(user.Password) {
		return nil, apperror.IncorrectCredentials
	}

	return stored, nil
}

func (us *UserService) SignUp(user *dto.SignUp) (*entities.User, error) {
	stored, err := us.repo.GetUserByEmail(user.Email)
	if err != nil && err != apperror.UserNotFound {
		return nil, err
	}

	if stored != nil {
		return nil, apperror.UserAlreadyExists
	}

	user.Password = utils.Hash(user.Password)
	toCreate := &entities.User{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	stored, err = us.repo.CreateUser(toCreate)
	if err != nil {
		return nil, err
	}
	return stored, nil
}

func (us *UserService) GetProfileByEmail(email string) (*entities.User, error) {
	user, err := us.repo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetProfileByLink(email string, link string) (*entities.User, error) {
	// TODO сравить email с email в найденой моели по линку и если не совпадает
	// TODO - запросить инфу о допутсимых действиях
	user, err := us.getUser(link)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) EditProfile(email string, profile *dto.EditProfile) error {
	_, err := us.repo.UpdateUser(email, profile)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Subscribe(subEmail, followLink string) error {
	reqTime := time.Now().String()
	_, err := us.repo.Subscribe(subEmail, followLink, reqTime)
	if err != nil {
		// проверить на повторную подписку
		return err
	}

	return nil
}

func (us *UserService) Unsubscribe(subEmail, followLink string) error {
	_, err := us.repo.Unsubscribe(subEmail, followLink)
	if err != nil {
		// проверить на повторную отписку
		return err
	}
	return nil
}

func (us *UserService) Reject(rejectEmail, followerLink string) error {
	err := us.repo.RejectFriendRequest(rejectEmail, followerLink)
	if err != nil {
		// валидация
		return err
	}
	return nil
}

func (us *UserService) GetFriendsByEmail(email string, limit, offset int) ([]*entities.User, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	users, err := us.repo.GetFriends(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetSubscribesByEmail(email string, limit, offset int) ([]*entities.User, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	users, err := us.repo.GetSubscribes(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetSubscribersByEmail(email string, limit, offset int) ([]*entities.User, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	users, err := us.repo.GetSubscribes(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetFriendsByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error) {
	user, err := us.getUser(targetLink)

	if err != nil {
		return nil, err
	}
	// TODO сравить requestEmail с email в найденой моели по линку и если не совпадает - запросить инфу о допутсимых действиях
	users, err := us.repo.GetFriends(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetSubscribesByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error) {
	user, err := us.getUser(targetLink)
	if err != nil {
		return nil, err
	}
	// TODO сравить requestEmail с email в найденой моели по линку и если не совпадает - запросить инфу о допутсимых действиях
	users, err := us.repo.GetSubscribes(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetSubscribersByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error) {
	user, err := us.getUser(targetLink)
	if err != nil {
		return nil, err
	}
	// TODO сравить requestEmail с email в найденой моели по линку и если не совпадает - запросить инфу о допутсимых действиях
	users, err := us.repo.GetSubscribers(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) getUser(link string) (*entities.User, error) {
	if strings.HasPrefix(link, "id") {
		id, err := strconv.Atoi(strings.TrimPrefix(link, "id"))
		if err != nil {
			return nil, apperror.BadRequest
		}
		return us.repo.GetUserById(uint(id))
	}

	return us.repo.GetUserByLink(link)
}
