package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository"
	"sort"
	"time"
)

type FeedService struct {
	repository repository.FeedRepository
}

func NewFeedService(feedRepository repository.FeedRepository) *FeedService {
	return &FeedService{
		repository: feedRepository,
	}
}

func (feed *FeedService) CollectPosts(user *entities.User, topPostDateTime time.Time, batchSize int) ([]entities.Post, error) {
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
		return posts[i].Date.Before(posts[j].Date)
	})

	if len(posts) >= batchSize {
		return posts[:batchSize], nil
	}

	return posts[:len(posts)], nil
}
