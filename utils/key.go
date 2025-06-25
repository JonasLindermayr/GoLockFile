package utils

import "crypto/sha256"

func passwordToKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}
