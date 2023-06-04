package service

import (
	"fmt"
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
	filename := strings.TrimPrefix(strings.TrimSuffix(url, "&type=img"), "static-service/download?name=")
	f, err := os.Open("static/files/img/" + filename)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		fmt.Println(err)
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

	if len(colors) == 1 {
		colors = append(colors, colors[0])
	}
	result := fmt.Sprintf("#%s #%s", colors[0].AsString(), colors[1].AsString())

	return result, nil
}
