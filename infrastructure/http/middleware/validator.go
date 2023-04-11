package middleware

import (
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/pkg/helpers"
	"net/http"
)

func UserRequestValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.UserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid request body. Please check your request body and try again",
			})
			return
		}

		if err := helpers.Validate(request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid request body. Please check your request body and try again",
			})
			return
		}

		c.Set("payload", request)
		c.Next()
	}
}

func AuthRequestValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.AuthRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid request body. Please check your request body and try again",
			})
			return
		}

		if err := helpers.Validate(request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid request body. Please check your request body and try again",
			})
			return
		}

		c.Set("payload", request)
		c.Next()
	}
}

func PhotoRequestValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.PhotoRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid request body. Please check your request body and try again",
			})
			return
		}

		if err := helpers.Validate(request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid request body. Please check your request body and try again",
			})
			return
		}

		c.Set("payload", request)
		c.Next()
	}
}
