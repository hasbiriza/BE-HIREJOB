package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secretKey, email string, role string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RefreshToken(secretKey, email string, role string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims for refresh token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	refreshToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
