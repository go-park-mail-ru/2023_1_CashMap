package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"errors"
)

type PostService struct {
	repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) usecase.PostUsecase {
	return &PostService{
		postRepository,
	}
}

func (service *PostService) GetPostById(email string, postDTO *dto.PostGetByID) (*entities.Post, error) {
	hasAccess, err := service.PostRepository.CheckReadAccess(email)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		return nil, errors.New("access to post denied")
	}

	post, err := service.SelectPostById(postDTO.PostID, email)
	if err != nil {
		return nil, err
	}

	owner, community, err := service.PostRepository.GetPostSenderInfo(post.ID)
	if err != nil {
		return nil, err
	}

	post.OwnerInfo = owner
	post.CommunityInfo = community

	return post, nil
}

func (service *PostService) GetPostsByCommunityLink(email string, dto *dto.PostsGetByLink) ([]*entities.Post, error) {
	hasAccess, err := service.PostRepository.CheckReadAccess(email)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		return nil, errors.New("access to posts denied")
	}

	posts, err := service.SelectPostsByCommunityLink(dto, email)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		owner, community, err := service.PostRepository.GetPostSenderInfo(post.ID)
		if err != nil {
			return nil, err
		}

		post.OwnerInfo = owner
		post.CommunityInfo = community
	}

	return posts, nil
}

func (service *PostService) GetPostsByUserLink(email string, dto *dto.PostsGetByLink) ([]*entities.Post, error) {
	hasAccess, err := service.PostRepository.CheckReadAccess(email)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		return nil, errors.New("access to posts denied")
	}

	if dto.LastPostDate == "" {
		dto.LastPostDate = "0"
	}

	posts, err := service.PostRepository.SelectPostsByUserLink(dto, email)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		owner, community, err := service.PostRepository.GetPostSenderInfo(post.ID)
		if err != nil {
			return nil, err
		}

		post.OwnerInfo = owner
		post.CommunityInfo = community
	}

	return posts, nil
}

func (service *PostService) CreatePost(email string, dto *dto.PostCreate) (*entities.Post, error) {
	if dto.Text == "" && dto.Attachments == nil {
		return nil, errors.New("empty input data")
	}

	if dto.CommunityLink != nil && dto.OwnerLink != nil {
		return nil, errors.New("too many data (community_link and owner_link can't come together)")
	}
	if dto.CommunityLink == nil && dto.OwnerLink == nil {
		return nil, errors.New("now enough input data")
	}

	hasAccess, err := service.PostRepository.CheckWriteAccess(email, dto)
	if err != nil {
		return nil, err
	}
	if !hasAccess {
		return nil, errors.New("access to posts denied")
	}
	dto = utils.Escaping(dto)
	postID, err := service.PostRepository.CreatePost(email, dto)
	if err != nil {
		return nil, err
	}
	return service.PostRepository.SelectPostById(postID, email)
}

func (service *PostService) UpdatePost(email string, dto *dto.PostUpdate) error {
	if dto.PostID == nil {
		return errors.New("post_id is required field")
	}

	err := service.PostRepository.UpdatePost(email, dto)
	if err != nil {
		return err
	}

	return nil
}

func (service *PostService) DeletePost(email string, dto *dto.PostDelete) error {
	err := service.PostRepository.DeletePost(email, dto)
	if err != nil {
		return err
	}

	return nil
}

func (service *PostService) LikePost(email string, dto *dto.LikeDTO) (int, error) {
	if dto.PostID == nil {
		return 0, apperror.NewServerError(apperror.BadRequest, errors.New("post_id can't be null"))
	}

	err := service.PostRepository.SetLike(email, *dto.PostID)
	if err != nil {
		return 0, err
	}

	return service.PostRepository.GetLikesAmount(email, *dto.PostID)
}

func (service *PostService) CancelLike(email string, dto *dto.LikeDTO) error {
	if dto.PostID == nil {
		return apperror.NewServerError(apperror.BadRequest, errors.New("post_id can't be null"))
	}

	err := service.PostRepository.CancelLike(email, *dto.PostID)
	if err != nil {
		return err
	}

	return nil
}

func (service *PostService) Repost() {
	//TODO: рк3
	panic("implement me")
}
