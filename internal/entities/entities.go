package entities

import "time"

type User struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Avatar     string    `json:"avatar"`
	Status     string    `json:"status"`
	Work       string    `json:"work"`
	Education  string    `json:"education"`
	BirthDate  time.Time `json:"birth_date"`
	DateJoined time.Time `json:"date_joined"`
	Friends    []User    `json:"friends"`
	Photos     []string  `json:"photos"`
	Groups     []Group   `json:"groups"`
	Posts      []Post    `json:"posts"`
}

type Group struct {
	ID           uint
	Title        string
	HeaderImage  string
	MembersCount int
	Owners       []User
	Posts        []Post
	// TODO: доделать поля модели для будущих потребностей
}

type Like struct {
	ID     uint
	PostID uint
	Sender User
}

type Comment struct {
	Text    string
	PostID  uint
	ReplyTo uint // id коммента в посте, к которому сделан коммент. null, если коммент верхнего уровня
}

type Post struct {
	ID          uint
	Text        string
	Attachments []string
	Likes       []Like
	Comments    []Comment
	Date        time.Time
}
