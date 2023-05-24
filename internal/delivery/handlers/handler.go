package handlers

type Handler struct {
	*UserHandler
	*FeedHandler
	*PostHandler
	*MessageHandler
	*GroupHandler
	*Sticker
	*CommentHandler
}

func NewHandler(userHandler *UserHandler,
	feedHandler *FeedHandler,
	postHandler *PostHandler,
	msgHandler *MessageHandler,
	groupHandler *GroupHandler,
	stickerHandler *Sticker,
	commentHandler *CommentHandler) Handler {

	return Handler{
		userHandler,
		feedHandler,
		postHandler,
		msgHandler,
		groupHandler,
		stickerHandler,
		commentHandler,
	}
}
