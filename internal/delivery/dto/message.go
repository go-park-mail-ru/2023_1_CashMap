package dto

type NewMessageDTO struct {
	UserId      uint     `json:"-"`
	ChatId      uint     `json:"chat_id"`
	ContentType string   `json:"message_content_type"`
	StickerID   *uint    `json:"sticker_id"`
	Text        string   `json:"text_content"`
	ReplyTo     *uint    `json:"reply_to"`
	Attachments []string `json:"attachments"`
}

type GetMessagesDTO struct {
	ChatID          *uint   `form:"chat_id" json:"chat_id" valid:"required"`
	BatchSize       *uint   `form:"batch_size" json:"batch_size"`
	LastMessageDate *string `form:"last_msg_date" json:"last_msg_date"`
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
