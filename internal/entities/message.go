package entities

//go:generate easyjson --all message.go

type Message struct {
	Id          *uint     `json:"id" db:"id"`
	UserId      *uint     `json:"-" db:"user_id"`
	SenderInfo  *UserInfo `json:"sender_info" db:"sender_info"`
	ChatId      *uint     `json:"chat_id" db:"chat_id"`
	ContentType *string   `json:"message_content_type" db:"message_content_type"`
	StickerId   *uint     `json:"-" db:"sticker_id"`
	Sticker     *Sticker  `json:"sticker"`
	Text        *string   `json:"text_content" db:"text_content"`
	CreatedAt   *string   `json:"creation_date" db:"creation_date"`
	ChangedAt   *string   `json:"change_date" db:"change_date"`
	ReplyTo     *uint     `json:"reply_to" db:"reply_to"`
	IsDeleted   *bool     `json:"is_deleted" db:"is_deleted"`
	Attachments []string  `json:"attachments" db:"attachments"`
}
