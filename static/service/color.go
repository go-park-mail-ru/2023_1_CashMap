package service

import (
	"depeche/static/repository"
	"github.com/EdlinOrg/prominentcolor"
	"image"
	"os"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type ColorUsecase interface {
	AverageColor(url string) (string, error)
}

type ColorService struct{}

func NewColorService() ColorUsecase {
	return &ColorService{}
}

func (c *ColorService) AverageColor(url string) (string, error) {
	filename := strings.TrimPrefix(strings.TrimSuffix(url, "&type=img"), "/static-service/download?name=")
	f, err := os.Open(repository.IMG_STATIC_PATH + filename)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}
	colors, err := prominentcolor.Kmeans(img)
	if err != nil {
		return "", err
	}

	return colors[0].AsString(), nil
}
