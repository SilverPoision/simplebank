package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passowrd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}

	return string(hash), nil
}

func ComparePassword(passowrd string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passowrd))
}
