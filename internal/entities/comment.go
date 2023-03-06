package entities

import "time"

type Comment struct {
	ID      uint      `json:"id"`
	Sender  string    `json:"sender_name"`
	Date    time.Time `json:"date"`
	Text    string    `json:"text"`
	PostID  uint      `json:"post_id"`
	ReplyTo uint      `json:"reply_to"`
	// id коммента в посте, к которому сделан коммент. null, если коммент верхнего уровня
}
