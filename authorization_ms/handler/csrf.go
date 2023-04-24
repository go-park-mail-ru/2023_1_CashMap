package handler

import (
	"context"
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/authEntities"
	"depeche/authorization_ms/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CSRF struct {
	service service.CSRFUsecase
	api.UnimplementedCSRFServiceServer
}

func NewCSRFHandler(service service.CSRFUsecase) api.CSRFServiceServer {
	return &CSRF{
		service: service,
	}
}

func (c *CSRF) CreateCSRFToken(ctx context.Context, email *api.EmailCSRF) (*api.TokenCSRF, error) {
	token, err := c.service.CreateCSRFToken(email.Email)
	if err != nil {
		return nil, err
	}
	return &api.TokenCSRF{Token: token}, nil
}

func (c *CSRF) InvalidateCSRFToken(ctx context.Context, csrf *api.CSRF) (*emptypb.Empty, error) {
	creds := &authEntities.CSRF{
		Email: csrf.Email,
		Token: csrf.Token,
	}
	err := c.service.InvalidateCSRFToken(creds)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (c *CSRF) ValidateCSRFToken(ctx context.Context, csrf *api.CSRF) (*api.Valid, error) {
	creds := &authEntities.CSRF{
		Email: csrf.Email,
		Token: csrf.Token,
	}
	valid, err := c.service.ValidateCSRFToken(creds)
	if err != nil {
		return nil, err
	}
	return &api.Valid{Valid: valid}, nil
}
