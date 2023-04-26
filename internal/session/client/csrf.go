package client

import (
	"context"
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/authEntities"
	"depeche/authorization_ms/service"
	"depeche/pkg/apperror"
)

type CSRFService struct {
	client api.CSRFServiceClient
	ctx    context.Context
}

func NewCSRFService(client api.CSRFServiceClient) service.CSRFUsecase {
	return &CSRFService{
		client: client,
		ctx:    context.Background(),
	}
}

func (c *CSRFService) CreateCSRFToken(email string) (string, error) {
	token, err := c.client.CreateCSRFToken(c.ctx, &api.EmailCSRF{Email: email})
	if err != nil {
		return "", apperror.NewServerError(apperror.NoAuth, err)
	}
	return token.Token, nil
}

func (c *CSRFService) InvalidateCSRFToken(csrf *authEntities.CSRF) error {
	_, err := c.client.InvalidateCSRFToken(c.ctx, &api.CSRF{Email: csrf.Email, Token: csrf.Token})
	if err != nil {
		return apperror.NewServerError(apperror.NoAuth, err)
	}
	return nil
}

func (c *CSRFService) ValidateCSRFToken(csrf *authEntities.CSRF) (bool, error) {
	valid, err := c.client.ValidateCSRFToken(c.ctx, &api.CSRF{Email: csrf.Email, Token: csrf.Token})
	if err != nil {
		return valid.Valid, apperror.NewServerError(apperror.NoAuth, err)
	}
	return valid.Valid, nil
}
