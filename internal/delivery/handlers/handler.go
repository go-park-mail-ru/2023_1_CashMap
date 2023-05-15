package handlers

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*MessageHandler
	*GroupHandler
}

func NewHandler(userHandler *UserHandler, feedHandler *FeedHandler, postHandler *PostHandler, msgHandler *MessageHandler, groupHandler *GroupHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		msgHandler,
		groupHandler,
	}
}
