package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"testing"
)

func TestFeedService_CollectPosts(t *testing.T) {
	var tests []struct {
		Name      string
		InputData struct {
			Email string
			Dto *dto.FeedDTO
		}
		ExpectedOutput []*entities.Post
		ExpectedErr error
	}{
		{

		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
