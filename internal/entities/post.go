package entities

import "time"

// Post entity info
//	@Description	All post information
type Post struct {
	ID          uint
	SenderName  string    `json:"sender_name"`
	Text        string    `json:"text"`
	Images      []string  `json:"images"`
	Attachments []string  `json:"attachments"`
	Likes       int       `json:"likes"`
	Comments    []Comment `json:"comments"`
	Date        time.Time `json:"date"`
}
