package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"fmt"
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
	// TODO generate link
	stored.Link = fmt.Sprintf("%d", stored.ID)
	return us.repo.UpdateUser(stored)
}

func (us *UserService) GetProfileByEmail(email string) (*dto.Profile, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	profile := &dto.Profile{
		Link:      user.Link,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		Status:    user.Status,
		Bio:       user.Bio,
		BirthDate: user.BirthDay,
	}
	return profile, nil
}

func (us *UserService) GetProfileByLink(link string) (*dto.Profile, error) {
	// TODO проверить есть ли доступ ко всем полям (закрытый аккаунт)
	// TODO добавить параметр email автора запроса
	user, err := us.repo.GetUserByLink(link)
	if err != nil {
		return nil, err
	}
	profile := &dto.Profile{
		Link:      user.Link,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		Status:    user.Status,
		Bio:       user.Bio,
		BirthDate: user.BirthDay,
	}
	return profile, nil
}

func (us *UserService) Subscribe(subEmail, followLink string) error {
	// подписаться (запись в user_subscriber)
	reqTime := time.Now().String()
	_, err := us.repo.Subscribe(subEmail, followLink, reqTime)
	if err != nil {
		// проверить на повторную подписку
		return err
	}

	return nil
}

func (us *UserService) Unsubscribe(subEmail, followLink string) error {

	// отписаться (запись в user_subscriber)
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
	user, err := us.repo.GetUserByLink(targetLink)
	if err != nil {
		return nil, err
	}
	// TODO добавить проверки разрешения requestEmail на просмотр друзей targetLink
	users, err := us.repo.GetFriends(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetSubscribesByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error) {
	user, err := us.repo.GetUserByLink(targetLink)
	if err != nil {
		return nil, err
	}
	// TODO добавить проверки разрешения requestEmail на просмотр подписок targetLink
	users, err := us.repo.GetSubscribes(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetSubscribersByLink(requestEmail, targetLink string, limit, offset int) ([]*entities.User, error) {
	user, err := us.repo.GetUserByLink(targetLink)
	if err != nil {
		return nil, err
	}
	// TODO добавить проверки разрешения requestEmail на просмотр подписчиков targetLink
	users, err := us.repo.GetSubscribers(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}
