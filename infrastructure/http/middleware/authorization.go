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

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := db.GetDBInstance()
		claim := c.MustGet("claim").(jwt.MapClaims)

		uid := int(claim["id"].(float64))
		smid, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid social media id. Social media id must be an integer",
			})
			return
		}

		var socialMedia entities.SocialMedia
		err = database.Select("user_id").First(&socialMedia, smid).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(http.StatusNotFound, dto.ApiResponse{
					Code:    http.StatusNotFound,
					Status:  "NOT_FOUND",
					Message: "The requested social media does not exist",
				})
				return
			}

			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
				Code:    http.StatusInternalServerError,
				Status:  "INTERNAL_SERVER_ERROR",
				Message: "An error occurred while processing your request. Please try again later",
			})
			return
		}

		if socialMedia.UserID != uid {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.ApiResponse{
				Code:    http.StatusForbidden,
				Status:  "FORBIDDEN",
				Message: "You are not authorized to perform this action",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := db.GetDBInstance()
		claim := c.MustGet("claim").(jwt.MapClaims)

		uid := int(claim["id"].(float64))
		cid, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid comment id. Comment id must be an integer",
			})
			return
		}

		var comment entities.Comment
		err = database.Select("user_id").First(&comment, cid).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(http.StatusNotFound, dto.ApiResponse{
					Code:    http.StatusNotFound,
					Status:  "NOT_FOUND",
					Message: "The requested comment does not exist",
				})
				return
			}

			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
				Code:    http.StatusInternalServerError,
				Status:  "INTERNAL_SERVER_ERROR",
				Message: "An error occurred while processing your request. Please try again later",
			})
			return
		}

		if comment.UserID != uid {
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

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := db.GetDBInstance()
		claim := c.MustGet("claim").(jwt.MapClaims)

		uid := int(claim["id"].(float64))
		pid, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "Invalid photo id. Photo id must be an integer",
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
					Message: "The requested photo does not exist",
				})
				return
			}

			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
				Code:    http.StatusInternalServerError,
				Status:  "INTERNAL_SERVER_ERROR",
				Message: "An error occurred while processing your request. Please try again later",
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
