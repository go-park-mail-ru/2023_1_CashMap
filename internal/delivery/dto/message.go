package dto

type NewMessage struct {
	UserId      uint   `json:"-"`
	Link        string `json:"user_link"`
	ChatId      uint   `json:"chat_id"`
	ContentType string `json:"message_content_type"`
	Text        string `json:"text_content"`
	ReplyTo     *uint  `json:"reply_to"`
}

type GetMessagesDTO struct {
	ChatID       *uint   `form:"chat_id" json:"chat_id" valid:"required"`
	BatchSize    *uint   `form:"batch_size" json:"batch_size"`
	LastPostDate *string `form:"last_post_date" json:"last_post_date"`
}

type CreateChatDTO struct {
	UserLinks []string `form:"user_links" json:"user_links" valid:"required"`
}

type GetChatsDTO struct {
	Offset    *uint `form:"offset" json:"offset"`
	BatchSize *uint `form:"batch_size" json:"batch_size" valid:"required"`
}

type HasDialogDTO struct {
	UserLink *string `form:"user_link" json:"user_link" valid:"required"`
}
