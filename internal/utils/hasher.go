package utils

import (
	"crypto/sha1"
	"fmt"
)

func Hash(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	passwordHash := fmt.Sprintf("%x", hasher.Sum(nil)) // кодируем сырые байты в Base64
	return passwordHash
}
