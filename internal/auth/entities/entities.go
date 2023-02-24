package entities

import (
	"github.com/Songmu/go-httpdate"
	"time"
)

type User struct {
	Credentials
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birth_date"`
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Session struct {
	Login          string `json:"login"`
	SessionId      string `json:"password"`
	ExpirationTime string `json:"expiration_time"`
	// TODO: надо как-то учесть часовой пояс....
}

func (session *Session) isExpired() (bool, error) {
	// захардкодил часовой пояс
	sessionTime, err := httpdate.Str2Time(session.ExpirationTime, time.Local)
	if err != nil {
		return false, err
	}

	if time.Now().Before(sessionTime) {
		return true, nil
	}

	return false, nil
}
