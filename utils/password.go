package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwordStr string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash the password : %w", err)
	}
	return string(hashedPass), nil
}
func CheckPassword(password, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
