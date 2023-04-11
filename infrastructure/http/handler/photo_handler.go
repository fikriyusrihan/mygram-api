package handler

import (
	"github.com/gin-gonic/gin"
	"my-gram/controllers"
)

func PostPhoto(ctr controllers.AppController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctr.HandleCreatePhoto(ctx)
	}
}

func PutPhoto(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleUpdatePhoto(c)
	}
}

func DeletePhoto(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleDeletePhoto(c)
	}
}

func GetPhotos(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleGetPhotos(c)
	}
}

func GetPhotoByID(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleGetPhotoByID(c)
	}
}
