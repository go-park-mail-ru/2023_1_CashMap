package entities

import "time"

type User struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Avatar     string    `json:"avatar"`
	Status     string    `json:"status"`
	Work       string    `json:"work"`
	Education  string    `json:"education"`
	BirthDate  time.Time `json:"birth_date"`
	DateJoined time.Time `json:"date_joined"`
	Friends    []User    `json:"friends"`
	Photos     []string  `json:"photos"`
}
