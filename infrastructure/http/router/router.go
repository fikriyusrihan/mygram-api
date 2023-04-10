package router

import (
	"github.com/gin-gonic/gin"
	"my-gram/controllers"
	"my-gram/infrastructure/http/handler"
	"my-gram/infrastructure/http/middleware"
)

func NewRouter(ctr controllers.AppController) *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.POST("/login", middleware.AuthRequestValidator(), handler.PostUserLogin(ctr))
		users.POST("/register", middleware.UserRequestValidator(), handler.PostUserRegister(ctr))
	}

	return router
}
