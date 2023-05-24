package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
	utils2 "depeche/internal/usecase/utils"
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

	err = service.AddPostData(post)
	if err != nil {
		return nil, err
	}
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

	if dto.LastPostDate == "" {
		dto.LastPostDate = utils2.OLDEST_DATE
	}

	posts, err := service.SelectPostsByCommunityLink(dto, email)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		err = service.AddPostData(post)
		if err != nil {
			return nil, err
		}
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
		dto.LastPostDate = utils2.OLDEST_DATE
	}

	posts, err := service.PostRepository.SelectPostsByUserLink(dto, email)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		err = service.AddPostData(post)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (service *PostService) CreatePost(email string, dto *dto.PostCreate) (*entities.Post, error) {
	if dto.Text == "" && dto.Attachments == nil {
		return nil, apperror.NewBadRequest()
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
	if dto.Attachments != nil && len(dto.Attachments) != 0 {
		if len(dto.Attachments) > 10 {
			return nil, apperror.NewServerError(apperror.TooMuchAttachments, nil)
		}
		err = service.PostRepository.AddPostAttachments(postID, dto.Attachments)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
	}

	post, err := service.PostRepository.SelectPostById(postID, email)
	if err != nil {
		return nil, err
	}
	err = service.AddPostData(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (service *PostService) UpdatePost(email string, dto *dto.PostUpdate) error {
	if dto.PostID == nil {
		return errors.New("post_id is required field")
	}
	attachments, err := service.PostRepository.GetPostAttachments(*dto.PostID)
	if err != nil {
		return err
	}

	if len(attachments)+len(dto.Attachments.Added)-len(dto.Attachments.Deleted) > 10 {
		return apperror.NewServerError(apperror.TooMuchAttachments, nil)
	}
	err = service.PostRepository.UpdatePostAttachments(*dto.PostID, dto.Attachments)
	if err != nil {
		return err
	}
	err = service.PostRepository.UpdatePost(email, dto)
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

func (service *PostService) AddPostData(post *entities.Post) error {
	owner, community, err := service.PostRepository.GetPostSenderInfo(post.ID)
	if err != nil {
		return err
	}

	post.OwnerInfo = owner
	post.CommunityInfo = community

	attachments, err := service.PostRepository.GetPostAttachments(post.ID)
	if err != nil {
		return err
	}
	post.Attachments = attachments
	return nil
}
