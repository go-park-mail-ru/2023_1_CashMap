package handlers

import staticDelivery "depeche/internal/static/delivery"

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*staticDelivery.FileHandler
}

func NewHandler(userHandler *UserHandler, feedHandler *FeedHandler, postHandler *PostHandler, staticHandler *staticDelivery.FileHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		staticHandler,
	}
}
