package entities

type Chat struct {
	ChatID uint        `db:"chat_id" json:"chat_id"`
	Users  []*UserInfo `db:"members" json:"members"`
}
