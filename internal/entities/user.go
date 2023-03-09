package entities

import "time"

// User entity info
//	@Description	User account information
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
	// Friends    []User    `json:"friends"`
	// Photos     []string  `json:"photos"`
	// Groups     []Group   `json:"groups"`
	// Posts      []Post    `json:"posts"`
}
