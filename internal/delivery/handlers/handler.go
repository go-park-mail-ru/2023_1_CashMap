package handlers

import staticDelivery "depeche/internal/static/delivery"

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*MessageHandler
	*staticDelivery.FileHandler
	*GroupHandler
	*Sticker
}

func NewHandler(userHandler *UserHandler,
	feedHandler *FeedHandler,
	postHandler *PostHandler,
	staticHandler *staticDelivery.FileHandler,
	msgHandler *MessageHandler,
	groupHandler *GroupHandler,
	stickerHandler *Sticker) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		msgHandler,
		staticHandler,
		groupHandler,
		stickerHandler,
	}
}
