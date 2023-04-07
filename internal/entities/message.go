package entities

type Message struct {
	Id          uint   `db:"id"`
	Link        string `db:"link"`
	ChatId      uint   `db:"chat_id"`
	ContentType string `db:"message_content_type"`
	Text        string `db:"text_content"`
	CreatedAt   string `db:"creation_date"`
	ChangedAt   string `db:"change_date"`
	ReplyTo     uint   `db:"reply_to"`
	IsDeleted   bool   `db:"is_deleted"`
}
