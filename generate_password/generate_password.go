package generate_password

import (
	"crypto/rand"
	"strings"
)

// letters - список допустимых символов в пароле
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	// Gets crypto/rand byte sequence
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Chooses random letter n times and returns the password
	// Not exactly uniformly random but not so far of it
	pass := make([]string, n)
	for i, val := range b {
		randIdx := val % uint8(len(letters))
		pass[i] = string(letters[randIdx])
	}
	return strings.Join(pass, ""), nil
}
