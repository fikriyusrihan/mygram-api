package di

import (
	"my-gram/controllers"
	"my-gram/infrastructure/db"
	"my-gram/repositories"
	"my-gram/services"
)

func NewAppControllerInstance() (controllers.AppController, error) {
	database, err := db.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.NewUserRepository(database)
	photoRepository := repositories.NewPhotoRepository(database)
	commentRepository := repositories.NewCommentRepository(database)
	socialMediaRepository := repositories.NewSocialMediaRepository(database)

	authService := services.NewAuthService(userRepository)
	photoService := services.NewPhotoService(photoRepository)
	commentService := services.NewCommentService(commentRepository, photoRepository)
	socialMediaService := services.NewSocialMediaService(socialMediaRepository)

	authController := controllers.NewAuthController(authService)
	photoController := controllers.NewPhotoController(photoService)
	commentController := controllers.NewCommentController(commentService)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	appController := controllers.NewAppController(
		authController,
		photoController,
		commentController,
		socialMediaController,
	)

	return appController, nil
}
