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
