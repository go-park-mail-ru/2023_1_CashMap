package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"sort"
)

type FeedService struct {
	repository repository.FeedRepository
}

func NewFeedService(feedRepository repository.FeedRepository) usecase.Feed {
	return &FeedService{
		repository: feedRepository,
	}
}

func (feed *FeedService) CollectPosts(email string, dto *dto.FeedDTO) ([]*entities.Post, error) {
	friendsPosts, err := feed.repository.GetFriendsPosts(email, dto)
	if err != nil {
		return nil, err
	}

	groupsPosts, err := feed.repository.GetGroupsPosts(email, dto)
	if err != nil {
		return nil, err
	}

	var posts []*entities.Post
	if friendsPosts != nil {
		posts = append(posts, friendsPosts...)
	}
	if groupsPosts != nil {
		posts = append(posts, groupsPosts...)
	}

	sort.Slice(posts, func(i int, j int) bool {
		return posts[i].CreationDate < posts[j].CreationDate
	})

	if len(posts) >= int(dto.BatchSize) {
		return posts[:dto.BatchSize], nil
	}

	return posts, nil
}
