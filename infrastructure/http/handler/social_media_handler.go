package handler

import (
	"github.com/gin-gonic/gin"
	"my-gram/controllers"
)

func PostSocialMedia(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleCreateSocialMedia(c)
	}
}

func PutSocialMedia(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleUpdateSocialMedia(c)
	}
}

func DeleteSocialMedia(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleDeleteSocialMedia(c)
	}
}

func GetSocialMediaByID(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleGetSocialMediaByID(c)
	}
}

func GetSocialMediaByUserID(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleGetSocialMediaByUserID(c)
	}
}
