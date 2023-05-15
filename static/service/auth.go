package service

import (
	"context"
	"depeche/authorization_ms/api"
	"depeche/pkg/apperror"
)

type AuthUsecase interface {
	CheckSession(token string) error
}

type AuthService struct {
	client api.AuthServiceClient
	ctx    context.Context
}

func NewAuthService(client api.AuthServiceClient) *AuthService {
	ctx := context.Background()
	return &AuthService{
		ctx:    ctx,
		client: client,
	}
}

func (service *AuthService) CheckSession(token string) error {
	_, err := service.client.CheckSession(service.ctx, &api.Token{
		Token: token,
	})
	if err != nil {
		return apperror.NewServerError(apperror.NoAuth, err)
	}

	return nil
}
