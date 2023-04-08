package utils

import (
	"fmt"
	"time"
)

var format = "%d-%02d-%02dT%02d:%02d:%02d"

func CurrentTimeString() string {
	t := time.Now()
	return fmt.Sprint(format,
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
