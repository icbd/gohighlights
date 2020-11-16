package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateToken return 2*length string
func GenerateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
