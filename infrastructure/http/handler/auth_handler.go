package handler

import (
	"github.com/gin-gonic/gin"
	"my-gram/controllers"
)

func PostUserLogin(ctr controllers.AppController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctr.HandleUserLogin(ctx)
	}
}

func PostUserRegister(ctr controllers.AppController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctr.HandleUserRegister(ctx)
	}
}
