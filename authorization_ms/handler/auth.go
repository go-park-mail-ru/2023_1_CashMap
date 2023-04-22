package handler

import (
	"context"
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Auth struct {
	service service.Auth
	api.UnimplementedAuthServiceServer
}

func NewAuthHandler(service service.Auth) api.AuthServiceServer {
	return &Auth{
		service: service,
	}
}

func (a *Auth) Authenticate(ctx context.Context, email *api.Email) (*api.Token, error) {
	token, err := a.service.Authenticate(email.Email)
	if err != nil {
		return nil, err
	}
	return &api.Token{Token: token}, nil
}

func (a *Auth) LogOut(ctx context.Context, token *api.Token) (*emptypb.Empty, error) {
	err := a.service.LogOut(token.Token)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a *Auth) CheckSession(ctx context.Context, token *api.Token) (*api.Email, error) {
	session, err := a.service.CheckSession(token.Token)
	if err != nil {
		return nil, err
	}
	return &api.Email{Email: session.Email}, nil
}
