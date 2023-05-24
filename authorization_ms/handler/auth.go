package handler

import (
	"context"
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/metrics"
	"depeche/authorization_ms/service"
	"depeche/pkg/apperror"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
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
	timer := prometheus.NewTimer(metrics.DurationHistogram.WithLabelValues("Authenticate"))
	defer timer.ObserveDuration()

	token, err := a.service.Authenticate(email.Email)
	if err != nil {
		metrics.RequestCounter.WithLabelValues(strconv.Itoa(apperror.Errors[err].Code), "Authenticate").Inc()
		return nil, err
	}
	metrics.RequestCounter.WithLabelValues("200", "Authenticate").Inc()
	return &api.Token{Token: token}, nil
}

func (a *Auth) LogOut(ctx context.Context, token *api.Token) (*emptypb.Empty, error) {
	timer := prometheus.NewTimer(metrics.DurationHistogram.WithLabelValues("LogOut"))
	defer timer.ObserveDuration()

	err := a.service.LogOut(token.Token)
	if err != nil {
		metrics.RequestCounter.WithLabelValues(strconv.Itoa(apperror.Errors[err].Code), "LogOut").Inc()
		return nil, err
	}
	metrics.RequestCounter.WithLabelValues("200", "LogOut").Inc()
	return &emptypb.Empty{}, nil
}

func (a *Auth) CheckSession(ctx context.Context, token *api.Token) (*api.Email, error) {
	timer := prometheus.NewTimer(metrics.DurationHistogram.WithLabelValues("CheckSession"))
	defer timer.ObserveDuration()

	session, err := a.service.CheckSession(token.Token)
	if err != nil {
		metrics.RequestCounter.WithLabelValues(strconv.Itoa(apperror.Errors[err].Code), "CheckSession").Inc()
		return nil, err
	}
	metrics.RequestCounter.WithLabelValues("200", "CheckSession").Inc()
	return &api.Email{Email: session.Email}, nil
}
