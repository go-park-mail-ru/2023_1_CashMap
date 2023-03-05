package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository/interface"
	"sort"
	"time"
)

type Feed interface {
	CollectPosts(user *entities.User, lastPostDate time.Time, batchSize uint) ([]entities.Post, error)
}

type FeedService struct {
	repository _interface.FeedRepository
}

func NewFeedService(feedRepository _interface.FeedRepository) *FeedService {
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
		return posts[i].Date.After(posts[j].Date)
	})

	if len(posts) >= int(batchSize) {
		return posts[:batchSize], nil
	}

	return posts[:len(posts)], nil
}
