package service

import (
	"depeche/internal/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
)

type PostService struct {
	repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) usecase.PostUsecase {
	return &PostService{
		postRepository,
	}
}

func (service *PostService) GetPostById(dto dto.PostDTO) (*entities.Post, error) {
	//TODO: проверка на закрытостть сообщества

}

func (service *PostService) GetPostsByCommunityLink() {

}

func (service *PostService) GetPostsByUserLink() {

}

func (service *PostService) CreatePost() {

}

func (service *PostService) DeletePost() {

}

func (service *PostService) LikePost() {

}

func (service *PostService) CancelLike() {

}
