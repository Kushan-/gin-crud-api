package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("supersecret")

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(secretKey)
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // interface {} any return type
		// type check
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected Signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("couldnt parse the token")
	}

	tokenISValid := parsedToken.Valid

	if !tokenISValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	fmt.Println(userId, email)
	return userId, nil
}
