package entities

//go:generate easyjson --all like.go

type Like struct {
	PostID uint
	Sender User
}

type LikesAmount struct {
	LikesAmount int `json:"likes_amount" db:"likes_amount"`
}
