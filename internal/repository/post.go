package repository

import "depeche/internal/entities"

type PostRepository interface {
	SelectPostById(postId string) (*entities.Post, error)
	SelectPostsByCommunityLink(communityLink string, batchSize int, offset int) ([]*entities.Post, error)
	SelectPostsByUserLink(userLink string, batchSize int, offset int) ([]*entities.Post, error)
}
