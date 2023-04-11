package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"my-gram/domain/dto"
	"my-gram/domain/entities"
	"my-gram/infrastructure/db"
	"net/http"
	"strconv"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := db.GetDBInstance()
		claim := c.MustGet("claim").(jwt.MapClaims)

		uid := uint(claim["id"].(float64))
		pid, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid photo id",
			})
			return
		}

		var photo entities.Photo
		err = database.Select("user_id").First(&photo, pid).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(http.StatusNotFound, dto.ApiResponse{
					Code:    http.StatusNotFound,
					Status:  "NOT_FOUND",
					Message: "Photo not found",
				})
				return
			}

			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
				Code:    http.StatusInternalServerError,
				Status:  "INTERNAL_SERVER_ERROR",
				Message: "Something went wrong",
			})
			return
		}

		if photo.UserID != uid {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.ApiResponse{
				Code:    http.StatusForbidden,
				Status:  "FORBIDDEN",
				Message: "You are not authorized to access this resource",
			})
			return
		}

		c.Next()
	}
}
