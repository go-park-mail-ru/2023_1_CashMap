package session

import (
	"time"
)

type Session struct {
	Email     string
	ExpiresAt time.Time
}

func (s *Session) isExpired() bool {
	if time.Now().Before(s.ExpiresAt) {
		return true
	}
	return false
}
