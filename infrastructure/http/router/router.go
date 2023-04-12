package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"my-gram/controllers"
	_ "my-gram/docs"
	"my-gram/infrastructure/http/handler"
	"my-gram/infrastructure/http/middleware"
)

func NewRouter(ctr controllers.AppController) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.POST("/login", middleware.AuthRequestValidator(), handler.PostUserLogin(ctr))
		users.POST("/register", middleware.UserRequestValidator(), handler.PostUserRegister(ctr))

		authorizedUsers := users.Group("", middleware.Authentication())
		{
			authorizedUsers.POST("/:userId/social_medias", middleware.SocialMediaRequestValidator(), handler.PostSocialMedia(ctr))
			authorizedUsers.GET("/:userId/social_medias", handler.GetSocialMediaByUserID(ctr))
			authorizedUsers.GET("/:userId/social_medias/:socialMediaId", handler.GetSocialMediaByID(ctr))

			authorizedSM := authorizedUsers.Group("/:userId/social_medias/:socialMediaId", middleware.SocialMediaAuthorization())
			{
				authorizedSM.PUT("", middleware.SocialMediaRequestValidator(), handler.PutSocialMedia(ctr))
				authorizedSM.DELETE("", handler.DeleteSocialMedia(ctr))
			}
		}
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

		comments := photos.Group("/:photoId/comments")
		{
			comments.POST("", middleware.CommentRequestValidator(), handler.PostComment(ctr))
			comments.GET("", handler.GetComments(ctr))
			comments.GET("/:commentId", handler.GetCommentByID(ctr))

			authorizedComments := comments.Group("/:commentId", middleware.CommentAuthorization())
			{
				authorizedComments.PUT("", middleware.CommentRequestValidator(), handler.PutComment(ctr))
				authorizedComments.DELETE("", handler.DeleteComment(ctr))
			}
		}
	}

	return router
}
