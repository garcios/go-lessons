package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	byeErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return byeErr == nil
}
