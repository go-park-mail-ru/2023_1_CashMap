package dto

type NewMessage struct {
	UserId      uint   `json:"-"`
	Link        string `json:"link"`
	ChatId      uint   `json:"chat_id"`
	ContentType string `json:"message_content_type"`
	Text        string `json:"text_content"`
	ReplyTo     uint   `json:"reply_to"`
}
