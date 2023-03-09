package session

import (
	"time"
)

type Session struct {
	Email     string
	ExpiresAt time.Time
}

//nolint:unused
func (s *Session) isExpired() bool {
	return time.Now().Before(s.ExpiresAt)
}
