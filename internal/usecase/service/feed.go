package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
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
		// самая "старая" дата
		dto.LastPostDate = "0"
	}

	if dto.BatchSize == 0 {
		return nil, apperror.BadRequest
	}
	//friendsPosts, err := feed.repository.GetFriendsPosts(email, dto)
	//if err != nil {
	//	return nil, err
	//}
	//
	//groupsPosts, err := feed.repository.GetGroupsPosts(email, dto)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var posts []*entities.Post
	//if friendsPosts != nil {
	//	posts = append(posts, friendsPosts...)
	//}
	//if groupsPosts != nil {
	//	posts = append(posts, groupsPosts...)
	//}
	posts, err := feed.repository.GetFeedPosts(email, dto)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		// TODO: поменять энтити поста (засылать даныне автора поста а не овнера)
		post.OwnerInfo, post.CommunityInfo, err = feed.postRepo.GetPostSenderInfo(post.ID)
		if err != nil {
			return nil, err
		}
	}
	sort.Slice(posts, func(i int, j int) bool {
		return posts[j].CreationDate < posts[i].CreationDate
	})

	if len(posts) >= int(dto.BatchSize) {
		return posts[:dto.BatchSize], nil
	}

	return posts, nil
}
