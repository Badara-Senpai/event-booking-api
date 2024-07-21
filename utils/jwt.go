package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// this help to do a type check, like to verify if two variables for examples are the same
		// like here we are checking if token.Method is the same than *jwt.SigningMethodHMAC, here we just need if it's true or false
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not Parse Token")
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("Invalid Token")
	}

	// As the comment on top of the function
	// Here we are checking if parsed.Token is the same than jwt.MapClaims
	// Now Here whe naeed the claims, if it's true it will return parsedToken.Claims and fill the ok var with true or false

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid Token Claims")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
