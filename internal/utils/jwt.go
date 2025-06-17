package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokens(userID uint) (string, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	})

	fmt.Println("Signing token with claims:", token)
	fmt.Println("Signing with secret:", []byte(os.Getenv("JWT_SECRET")))

	tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	refreshStr, _ := refresh.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenStr, refreshStr
}

func RefreshAccessToken(refresh string) (string, error) {
	token, _ := jwt.Parse(refresh, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(float64)
		newToken, _ := GenerateTokens(uint(userID))
		return newToken, nil
	}
	return "", errors.New("invalid refresh token")
}
