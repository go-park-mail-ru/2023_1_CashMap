package entities

type Message struct {
	Id          *uint              `json:"id" db:"id"`
	UserId      *uint              `json:"-" db:"user_id"`
	Link        *string            `json:"link" db:"link"`
	SenderInfo  *MessageSenderInfo `json:"sender_info" db:"sender_info"`
	ChatId      *uint              `json:"chat_id" db:"chat_id"`
	ContentType *string            `json:"message_content_type" db:"message_content_type"`
	Text        *string            `json:"text_content" db:"text_content"`
	CreatedAt   *string            `json:"creation_date" db:"creation_date"`
	ChangedAt   *string            `json:"change_date" db:"change_date"`
	ReplyTo     *uint              `json:"reply_to" db:"reply_to"`
	IsDeleted   *bool              `json:"is_deleted" db:"is_deleted"`
}

type MessageSenderInfo struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Url       string `json:"url" db:"url"`
	Link      string `json:"link" db:"link"`
}
