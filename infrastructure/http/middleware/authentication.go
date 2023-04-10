package middleware

import (
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/pkg/helpers"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResponse{
				Code:    http.StatusUnauthorized,
				Status:  "UNAUTHORIZED",
				Message: "Invalid access token. Please login to get a new access token",
			})
			return
		}

		c.Set("claim", claim)
		c.Next()
	}
}
