package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secretKey")

// TODO: Encode/decode JWT

// Create a token for the given user ID
func makeToken(userId string) string {
	expirationTime := time.Now().Add(time.Hour)

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return ""
	}

	return signedToken
}

// Validate the token for the given user ID, returns an error if the token is invalid.
func verifyToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Extract and return claims if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["userId"].(string)
		if !ok {
			return "", errors.New("invalid token payload")
		}
		return userId, nil
	}

	return "", errors.New("invalid token")
}
