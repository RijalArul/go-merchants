package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - untuk membuat hash dari password plaintext
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash - untuk membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
