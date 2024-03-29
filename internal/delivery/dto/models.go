package dto

import (
	"depeche/internal/entities"
	"strconv"
)

//go:generate easyjson --all models.go

// [INCOMING]

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type EditProfile struct {
	Email            *string `json:"email"`
	NewPassword      *string `json:"password"`
	PreviousPassword *string `json:"prev_pass"`
	FirstName        *string `json:"first_name"`
	LastName         *string `json:"last_name"`
	Avatar           *string `json:"avatar_url"`
	Link             *string `json:"user_link"`
	Sex              *string `json:"sex"`
	Status           *string `json:"status"`
	Bio              *string `json:"bio"`
	Birthday         *string `json:"birthday"`
}

type Subscribes struct {
	Link string `json:"user_link"`
}

type GlobalSearchDTO struct {
	SearchQuery *string `form:"search_query"`
	BatchSize   *uint   `form:"batch_size"`
	Offset      *uint   `form:"offset"`
}

// [OUTGOING]

type Profile struct {
	Link       string `json:"user_link"   example:"id100500"`
	FirstName  string `json:"first_name"  example:"Василий"`
	LastName   string `json:"last_name"   example:"Петров"`
	Avatar     string `json:"avatar_url"  example:""`
	AvgColor   string `json:"avg_avatar_color"`
	Sex        string `json:"sex"         example:"male"`
	Status     string `json:"status"      example:"Текст статуса."`
	Bio        string `json:"bio"         example:"Текст с информацией о себе."`
	BirthDay   string `json:"birthday"    example:"30.04.2002"`
	DateJoined string `json:"date_joined" example:"10.02.2023"`
	LastActive string `json:"last_active" example:""`
	Private    bool   `json:"private"     example:"false"`
}

type UserStatus int

const (
	None = iota
	Friend
	Subscriber
	Subscribed
)

var StatusToString = map[UserStatus]string{
	None:       "none",
	Friend:     "friend",
	Subscriber: "subscriber",
	Subscribed: "subscribed",
}

func NewProfileFromUser(user *entities.User) *Profile {
	// TODO add fields to entity
	if user.Link == "" {
		user.Link = "id" + strconv.Itoa(int(user.ID))
	}
	return &Profile{
		Link:       user.Link,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Avatar:     user.Avatar,
		AvgColor:   user.AvgAvatarColor,
		Status:     user.Status,
		Bio:        user.Bio,
		BirthDay:   user.BirthDay,
		Private:    false,
		Sex:        user.Sex,
		LastActive: user.LastActive,
	}
}
