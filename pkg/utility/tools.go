package utility

import (
	"crypto/rand"
	"time"
)

// GenerateID generates a random string of the specified length.
func GenerateID(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// GetTimeNow returns the current time in UTC.
func GetTimeNow() time.Time {
	return time.Now().UTC()
}
