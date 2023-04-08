package entities

type Chat struct {
	ChatID    uint     `db:"chat_id" json:"chat_id"`
	UserLinks []string `db:"user_links" json:"user_links"`
}
