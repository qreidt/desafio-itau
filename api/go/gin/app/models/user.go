package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document" gorm:"uniqueIndex"`
	Password  string    `json:"-"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HashPassword(password string) (string, error) {
	bytes := []byte(password)

	// Hash the password with the default salt
	hashedBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword := string(hashedBytes)
	return hashedPassword, nil
}

func ComparePassword(hash string, passwordAttempt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordAttempt))
	return err == nil
}
