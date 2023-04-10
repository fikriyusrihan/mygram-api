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

	authService := services.NewAuthService(userRepository)

	authController := controllers.NewAuthController(authService)

	appController := controllers.NewAppController(authController)

	return appController, nil
}
