package handlers

import "depeche/internal/delivery"

type Handler struct {
	delivery.UserHandler
	delivery.FeedHandler
	delivery.PostHandler
}

func NewHandler(userHandler delivery.UserHandler, feedHandler delivery.FeedHandler, postHandler delivery.PostHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
	}
}
