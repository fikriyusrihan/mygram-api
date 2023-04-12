package controllers

type AppController interface {
	AuthController
	PhotoController
	CommentController
	SocialMediaController
}

type appController struct {
	AuthController
	PhotoController
	CommentController
	SocialMediaController
}

func NewAppController(
	authController AuthController,
	photoController PhotoController,
	commentController CommentController,
	socialMediaController SocialMediaController,
) AppController {
	return &appController{
		AuthController:        authController,
		PhotoController:       photoController,
		CommentController:     commentController,
		SocialMediaController: socialMediaController,
	}
}
