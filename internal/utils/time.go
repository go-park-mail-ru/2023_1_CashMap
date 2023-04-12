package utils

import (
	"time"
)

func CurrentTimeString() string {
	currentTime := time.Now().Format(time.RFC3339)
	return currentTime
}
