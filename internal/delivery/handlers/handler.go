package handlers

import staticDelivery "depeche/internal/static/delivery"

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*MessageHandler
	*staticDelivery.FileHandler
	*GroupHandler
}

func NewHandler(userHandler *UserHandler, feedHandler *FeedHandler, postHandler *PostHandler, staticHandler *staticDelivery.FileHandler, msgHandler *MessageHandler, groupHandler *GroupHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		msgHandler,
		staticHandler,
		groupHandler,
	}
}
