package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash returns the sha256 hash of the given string
func Hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func CompareHash(hash, s string) bool {
	return hash == Hash(s)
}
