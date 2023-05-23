package response

import "depeche/internal/entities"

type GetFeedResponse struct {
	Body GetFeedBody `json:"body"`
}

type GetFeedBody struct {
	Posts []*entities.Post `json:"posts"`
}
