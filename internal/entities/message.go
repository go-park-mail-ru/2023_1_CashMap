package entities

type Message struct {
	Id          uint   `db:"id" json:"id"`
	Link        string `db:"link" json:"link"`
	ChatId      uint   `db:"chat_id" json:"chat_id"`
	ContentType string `db:"message_content_type" json:"content_type"`
	Text        string `db:"text_content" json:"text"`
	CreatedAt   string `db:"creation_date" json:"creation_date"`
	ChangedAt   string `db:"change_date" json:"change_date"`
	ReplyTo     uint   `db:"reply_to" json:"reply_to"`
	IsDeleted   bool   `db:"is_deleted" json:"is_deleted"`
}
