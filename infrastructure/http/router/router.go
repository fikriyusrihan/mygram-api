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

	photos := router.Group("/photos", middleware.Authentication())
	{
		photos.POST("", middleware.PhotoRequestValidator(), handler.PostPhoto(ctr))
		photos.GET("", handler.GetPhotos(ctr))
		photos.GET("/:photoId", handler.GetPhotoByID(ctr))

		authorizedPhotos := photos.Group("/:photoId", middleware.PhotoAuthorization())
		{
			authorizedPhotos.PUT("", middleware.PhotoRequestValidator(), handler.PutPhoto(ctr))
			authorizedPhotos.DELETE("", handler.DeletePhoto(ctr))
		}
	}

	return router
}
