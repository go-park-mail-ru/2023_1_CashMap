package session

import "time"

type Session struct {
	Email     string
	ExpiresAt time.Time
}
