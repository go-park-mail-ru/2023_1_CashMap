package dto

type CreateCommentDTO struct {
	PostID  *uint   `json:"post_id"`
	ReplyTo *string `json:"reply_to"`
	Text    *string `json:"text"`
}

type EditCommentDTO struct {
	ID   *uint   `json:"id"`
	Text *string `json:"text"`
}

type GetCommentsDTO struct {
	ID              uint    `form:"id"`
	LastCommentDate *string `form:"last_comment_date"`
	Count           *string `form:"batch_size"`
}
