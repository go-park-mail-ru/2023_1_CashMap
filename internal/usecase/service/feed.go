package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"fmt"
	"sort"
)

type FeedService struct {
	repository repository.FeedRepository
	postRepo   repository.PostRepository
}

func NewFeedService(feedRepository repository.FeedRepository, postRepo repository.PostRepository) usecase.Feed {
	return &FeedService{
		repository: feedRepository,
		postRepo:   postRepo,
	}
}

func (feed *FeedService) CollectPosts(email string, dto *dto.FeedDTO) ([]*entities.Post, error) {
	if dto.LastPostDate == "" {
		dto.LastPostDate = "0"
	}
	fmt.Println(dto.BatchSize)
	if dto.BatchSize == 0 {
		return nil, apperror.BadRequest
	}
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
	for _, post := range posts {
		post.OwnerInfo, post.CommunityInfo, err = feed.postRepo.GetPostSenderInfo(post.ID)
		if err != nil {
			return nil, err
		}
	}
	sort.Slice(posts, func(i int, j int) bool {
		return posts[i].CreationDate < posts[j].CreationDate
	})

	if len(posts) >= int(dto.BatchSize) {
		return posts[:dto.BatchSize], nil
	}

	return posts, nil
}
