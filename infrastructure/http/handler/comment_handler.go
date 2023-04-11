package handler

import (
	"github.com/gin-gonic/gin"
	"my-gram/controllers"
)

func PostComment(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleCreateComment(c)
	}
}

func PutComment(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleUpdateComment(c)
	}
}

func DeleteComment(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleDeleteComment(c)
	}
}

func GetCommentByID(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleGetCommentByID(c)
	}
}

func GetComments(ctr controllers.AppController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctr.HandleGetComments(c)
	}
}
