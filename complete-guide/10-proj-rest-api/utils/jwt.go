package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	apperrors "github.com/oskiegarcia/event-booking/app-errors"
	"time"
)

const secretKey = "super-secret-key"

func GenerateToken(email, userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", apperrors.TokenGenerationError
	}

	return tokenString, nil
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return userID, nil
}
