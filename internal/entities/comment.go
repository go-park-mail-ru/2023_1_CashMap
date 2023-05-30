package entities

//go:generate easyjson --all comment.go

type Comment struct {
	ID            *uint              `json:"id" db:"id"`
	PostID        *uint              `json:"post_id" db:"post_id"`
	SenderInfo    *CommentSenderInfo `json:"sender_info" db:"sender_info"`
	ReplyReceiver *CommentSenderInfo `json:"reply_receiver_info" db:"reply_receiver_info"`
	Text          *string            `json:"text" db:"text_content"`
	CreationDate  *string            `json:"creation_date" db:"creation_date"`
	ChangeDate    *string            `json:"change_date" db:"change_date"`
	IsDeleted     *bool              `json:"is_deleted" db:"is_deleted"`
	IsAuthor      *bool              `json:"is_author" db:"is_author"`
}

type CommentSenderInfo struct {
	FirstName *string `json:"first_name" db:"first_name"`
	LastName  *string `json:"last_name" db:"last_name"`
	UserLink  *string `json:"user_link" db:"link"`
	AvatarUrl *string `json:"avatar_url" db:"avatar_url"`
}
