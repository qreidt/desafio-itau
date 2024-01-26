package util

import "golang.org/x/crypto/bcrypt"

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
