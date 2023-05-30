package response

import "depeche/internal/entities"

//go:generate easyjson --all feed.go

type GetFeedResponse struct {
	Body GetFeedBody `json:"body"`
}

type GetFeedBody struct {
	Posts []*entities.Post `json:"posts"`
}
