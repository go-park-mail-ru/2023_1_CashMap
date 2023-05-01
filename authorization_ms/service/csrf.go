package service

import (
	"depeche/authorization_ms/authEntities"
	"depeche/authorization_ms/repository"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type CSRFUsecase interface {
	CreateCSRFToken(email string) (string, error)
	InvalidateCSRFToken(csrf *authEntities.CSRF) error
	ValidateCSRFToken(csrf *authEntities.CSRF) (bool, error)
}

type CSRFService struct {
	repository repository.CSRFRepository
}

func NewCSRFService(repository repository.CSRFRepository) *CSRFService {
	return &CSRFService{repository}
}

func (service *CSRFService) CreateCSRFToken(email string) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", apperror.InternalServerError
	}

	csrf := &authEntities.CSRF{
		Email: email,
		Token: id.String(),
	}

	expirationTime := int((time.Hour * 24).Minutes())
	fmt.Println(expirationTime)

	err = service.repository.SaveCSRFToken(csrf, expirationTime)
	if err != nil {
		return "", err
	}

	return csrf.Token, nil
}

func (service *CSRFService) InvalidateCSRFToken(csrf *authEntities.CSRF) error {
	err := service.repository.DeleteCSRFToken(csrf)
	if err != nil {
		return err
	}

	return nil
}

func (service *CSRFService) ValidateCSRFToken(csrf *authEntities.CSRF) (bool, error) {
	success, err := service.repository.CheckCSRFToken(csrf)
	if err != nil {
		fmt.Println(err)
		return false, apperror.Forbidden
	}

	return success, nil
}
