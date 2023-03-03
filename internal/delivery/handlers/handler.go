package handlers

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
}

func NewHandler(userHandler *UserHandler, feedHandler *FeedHandler, postHandler *PostHandler) Handler {
	return Handler{
		userHandler,
		feedHandler,
		postHandler,
	}
}
