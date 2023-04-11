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

	authService := services.NewAuthService(userRepository)
	photoService := services.NewPhotoService(photoRepository)

	authController := controllers.NewAuthController(authService)
	photoController := controllers.NewPhotoController(photoService)

	appController := controllers.NewAppController(
		authController,
		photoController,
	)

	return appController, nil
}
