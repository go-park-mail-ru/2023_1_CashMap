package dto

import (
	"depeche/internal/entities"
	"strconv"
)

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
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Link      *string `json:"link"`
	Sex       *string `json:"sex"`
	Status    *string `json:"status"`
	Bio       *string `json:"bio"`
	Birthday  *string `json:"birthday"`
}

type Subscribes struct {
	Link string `json:"link"`
}

// [OUTGOING]

type Profile struct {
	Email      string `json:"email"       example:"example@mail.ru"`
	Link       string `json:"link"        example:"id100500"`
	FirstName  string `json:"first_name"  example:"Василий"`
	LastName   string `json:"last_name"   example:"Петров"`
	Avatar     string `json:"avatar"      example:""`
	Sex        string `json:"sex"         example:"male"`
	Status     string `json:"status"      example:"Текст статуса."`
	Bio        string `json:"bio"         example:"Текст с информацией о себе."`
	BirthDay   string `json:"birthday"    example:"30.04.2002"`
	DateJoined string `json:"date_joined" example:"10.02.2023"`
	LastActive string `json:"last_active" example:""`
	Private    bool   `json:"private"     example:"false"`
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
		Status:     user.Status,
		Bio:        user.Bio,
		BirthDay:   user.BirthDay,
		Private:    false,
		Sex:        user.Sex,
		LastActive: "",
	}
}

func (si *SignIn) AuthEmail() string {
	return si.Email
}

func (su *SignUp) AuthEmail() string {
	return su.Email
}