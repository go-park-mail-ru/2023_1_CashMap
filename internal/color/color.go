package color

import (
	"context"
	"depeche/pkg/apperror"
	api "depeche/static/static_grpc"
)

type AvgColorUsecase interface {
	AverageColor(url string) (string, error)
}

type AvgColorService struct {
	client api.ColorServiceClient
	ctx    context.Context
}

func NewAvgColorService(client api.ColorServiceClient) AvgColorUsecase {
	return &AvgColorService{
		client: client,
		ctx:    context.Background(),
	}
}

func (a *AvgColorService) AverageColor(url string) (string, error) {
	color, err := a.client.AverageColor(a.ctx, &api.Url{
		Url: url,
	})
	if err != nil {
		return "", apperror.NewServerError(apperror.BadRequest, err)
	}
	return color.ColorStr, nil
}
