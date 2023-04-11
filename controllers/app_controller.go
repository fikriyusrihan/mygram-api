package controllers

type AppController interface {
	AuthController
	PhotoController
}

type appController struct {
	AuthController
	PhotoController
}

func NewAppController(
	authController AuthController,
	photoController PhotoController,
) AppController {
	return &appController{
		AuthController:  authController,
		PhotoController: photoController,
	}
}
