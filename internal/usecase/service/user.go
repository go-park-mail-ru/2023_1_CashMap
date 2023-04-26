package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
	"strconv"
	"strings"
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
	user = utils.Escaping(user)
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

func (us *UserService) GetAllUsers(email string, limit, offset int) ([]*entities.User, error) {
	users, err := us.repo.GetUsers(email, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) EditProfile(email string, profile *dto.EditProfile) error {

	if profile.NewPassword != nil {
		user, err := us.repo.GetUserByEmail(email)
		if err != nil {
			return err
		}
		if profile.PreviousPassword == nil {
			return apperror.Forbidden
		}
		if user.Password != utils.Hash(*profile.PreviousPassword) {
			return apperror.Forbidden
		}
		*profile.NewPassword = utils.Hash(*profile.NewPassword)
	}

	if profile.Avatar != nil {
		err := us.repo.UpdateAvatar(email, *profile.Avatar)
		if err != nil {
			return err
		}
		profile.Avatar = nil
	}

	// TODO validate errors
	if profile.Link != nil {
		exists, err := us.repo.CheckLinkExists(*profile.Link)
		if err != nil {
			return err
		}

		if exists {
			return apperror.UserAlreadyExists
		}
	}

	profile = utils.Escaping(profile)
	_, err := us.repo.UpdateUser(email, profile)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Subscribe(subEmail, followLink string) error {
	reqTime := utils.CurrentTimeString()
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
	users, err := us.repo.GetSubscribers(user, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetPendingRequestsByEmail(email string, limit, offset int) ([]*entities.User, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	users, err := us.repo.GetPendingFriendRequests(user, limit, offset)
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

func (us *UserService) GlobalSearch(dto *dto.GlobalSearchDTO) ([]*entities.UserInfo, error) {
	if dto.SearchQuery == nil || strings.TrimSpace(*dto.SearchQuery) == "" {
		return nil, apperror.NewServerError(apperror.BadRequest, errors.New("search query is required"))
	}

	if dto.BatchSize == nil {
		dto.BatchSize = new(uint)
		*dto.BatchSize = 0
	}

	if dto.Offset == nil {
		dto.Offset = new(uint)
		*dto.Offset = 0
	}

	return us.repo.SearchUserByName(dto)
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
