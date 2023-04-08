package handlers

import staticDelivery "depeche/internal/static/delivery"

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*MessageHandler
	*staticDelivery.FileHandler
}

func NewHandler(userHandler *UserHandler, feedHandler *FeedHandler, postHandler *PostHandler,
	messageHandler *MessageHandler, staticHandler *staticDelivery.FileHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		messageHandler,
		staticHandler,
	}
}
