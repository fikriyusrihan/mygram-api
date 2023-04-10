package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"my-gram/config"
	"my-gram/domain/entities"
	"time"
)

var secretKey = config.C.JWT.SignatureKey

func GenerateToken(user entities.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, err
}
