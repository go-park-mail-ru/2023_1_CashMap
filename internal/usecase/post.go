package usecase

type PostUsecase interface {
	GetPostById()
	GetPostsByCommunityLink()
	GetPostsByUserLink()

	CreatePost()

	DeletePost()

	LikePost()
	CancelLike()
	// TODO: Сделать обновление поста
	//UpdatePost()

	//AddComment() - в comment service
	//RemoveComment()
	// TODO: сделать центр уведомлений для комментариев и сообщений входящий (да и для вообще всего)

}
