package delivery

import (
	"context"
	"depeche/static/service"
	api "depeche/static/static_grpc"
)

type ColorHandler struct {
	colorService service.ColorUsecase
	api.UnimplementedColorServiceServer
}

func NewColorHandler(service service.ColorUsecase) *ColorHandler {
	return &ColorHandler{
		colorService: service,
	}
}

func (c *ColorHandler) AverageColor(ctx context.Context, in *api.Url) (*api.Color, error) {
	avgColorHex, err := c.colorService.AverageColor(in.Url)
	if err != nil {
		return nil, err
	}
	color := &api.Color{
		ColorStr: avgColorHex,
	}
	return color, nil
}
