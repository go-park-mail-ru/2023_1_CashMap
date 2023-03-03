package entities

import "time"

type Comment struct {
	ID      uint
	Sender  string
	Date    time.Time
	Text    string
	PostID  uint
	ReplyTo uint // id коммента в посте, к которому сделан коммент. null, если коммент верхнего уровня
}
