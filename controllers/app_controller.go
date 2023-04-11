package controllers

type AppController interface {
	AuthController
	PhotoController
	CommentController
}

type appController struct {
	AuthController
	PhotoController
	CommentController
}

func NewAppController(
	authController AuthController,
	photoController PhotoController,
	commentController CommentController,
) AppController {
	return &appController{
		AuthController:    authController,
		PhotoController:   photoController,
		CommentController: commentController,
	}
}
