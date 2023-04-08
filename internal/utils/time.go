package utils

import (
	"time"
)

func CurrentTimeString() string {
	currentTime := time.Now().String()
	return currentTime
}
