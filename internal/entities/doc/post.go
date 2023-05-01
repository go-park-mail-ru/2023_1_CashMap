package doc

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

// PostResponse entity info
//
//	@Description	All post information
type PostResponse struct {
	Body entities.Post `json:"body"`
}

// PostsResponse entity info
//
//	@Description	All post information
type PostsResponse struct {
	Body []entities.Post `json:"body"`
}

// PostDelete entity info
//
//	@Description	All post information
type PostDelete struct {
	Body dto.PostDelete `json:"body"`
}

type PostArray struct {
	Body []entities.Post `json:"body"`
}

type LikePost struct {
	Body entities.LikesAmount `json:"body"`
}
