package handlers

import staticDelivery "depeche/internal/static/delivery"

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*staticDelivery.FileHandler
	*MessageHandler
}

func NewHandler(userHandler *UserHandler, feedHandler *FeedHandler, postHandler *PostHandler, staticHandler *staticDelivery.FileHandler, msgHandler *MessageHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		staticHandler,
		msgHandler,
	}
}
