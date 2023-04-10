package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-gram/config"
	"my-gram/domain/entities"
	"strings"
	"time"
)

var secretKey = config.C.JWT.SignatureKey

func GenerateToken(user *entities.User) (string, error) {
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

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("please sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")

	isBearer := strings.HasPrefix(headerToken, "Bearer ")
	if !isBearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
