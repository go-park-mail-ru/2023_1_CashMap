package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"sort"
	"time"
)

type FeedService struct {
	repository repository.FeedRepository
}

func NewFeedService(feedRepository repository.FeedRepository) usecase.Feed {
	return &FeedService{
		repository: feedRepository,
	}
}

func (feed *FeedService) CollectPosts(user *entities.User, topPostDateTime time.Time, batchSize uint) ([]entities.Post, error) {
	friendsPosts, err := feed.repository.GetFriendsPosts(user, topPostDateTime, batchSize)
	if err != nil {
		return nil, err
	}

	groupsPosts, err := feed.repository.GetGroupsPosts(user, topPostDateTime, batchSize)
	if err != nil {
		return nil, err
	}

	var posts []entities.Post
	posts = append(posts, friendsPosts...)
	posts = append(posts, groupsPosts...)

	sort.Slice(posts, func(i int, j int) bool {
		return posts[i].CreationDate < posts[j].CreationDate
	})

	if len(posts) >= int(batchSize) {
		return posts[:batchSize], nil
	}

	return posts, nil
}
