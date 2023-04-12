package controllers

import (
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/services"
	"net/http"
)

type AuthController interface {
	HandleUserLogin(c *gin.Context)
	HandleUserRegister(c *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{authService}
}

// HandleUserLogin User login handler
// @Summary Authenticate User
// @Description Authenticate registered user with email and password
// @Tags Auth
// @Param Payload body dto.AuthRequest true "User Request Payload"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.AuthResponse}
// @Router /users/login [post]
func (ctr authController) HandleUserLogin(c *gin.Context) {
	payload := c.MustGet("payload").(dto.AuthRequest)
	response, err := ctr.authService.Login(&payload)
	if err != nil {
		c.AbortWithStatusJSON(err.Code(), dto.ApiResponse{
			Code:    err.Code(),
			Status:  err.Status(),
			Message: err.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login successful. Please use the token to access protected resources",
		Data:    response,
	})
}

// HandleUserRegister User register handler
// @Summary Register User
// @Description Register new user with username, email, age, and password
// @Tags Auth
// @Param Payload body dto.UserRequest true "User Request Payload"
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.ApiResponse{data=dto.UserResponse}
// @Router /users/register [post]
func (ctr authController) HandleUserRegister(c *gin.Context) {
	payload := c.MustGet("payload").(dto.UserRequest)
	response, err := ctr.authService.Register(&payload)
	if err != nil {
		c.AbortWithStatusJSON(err.Code(), dto.ApiResponse{
			Code:    err.Code(),
			Status:  err.Status(),
			Message: err.Message(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "User created successfully",
		Data:    response,
	})
}
