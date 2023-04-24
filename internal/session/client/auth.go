package client

import (
	"context"
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/authEntities"
	"depeche/authorization_ms/service"
	"depeche/pkg/apperror"
)

type AuthService struct {
	client api.AuthServiceClient
	ctx    context.Context
}

func NewAuthService(client api.AuthServiceClient) service.Auth {
	return &AuthService{
		client: client,
		ctx:    context.Background(),
	}
}

func (a *AuthService) Authenticate(email string) (string, error) {
	token, err := a.client.Authenticate(a.ctx, &api.Email{Email: email})
	if err != nil {
		return "", apperror.NewServerError(apperror.NoAuth, err)
	}
	return token.Token, nil
}

func (a *AuthService) LogOut(token string) error {
	_, err := a.client.LogOut(a.ctx, &api.Token{Token: token})
	if err != nil {
		return apperror.NewServerError(apperror.NoAuth, err)
	}
	return nil
}

func (a *AuthService) CheckSession(token string) (*authEntities.Session, error) {
	email, err := a.client.CheckSession(a.ctx, &api.Token{Token: token})
	if err != nil {
		return nil, apperror.NewServerError(apperror.NoAuth, err)
	}
	session := &authEntities.Session{Email: email.Email}
	return session, nil
}
