package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/services"
	"net/http"
)

type PhotoController interface {
	HandleCreatePhoto(c *gin.Context)
	HandleUpdatePhoto(c *gin.Context)
	HandleDeletePhoto(c *gin.Context)
	HandleGetPhotoByID(c *gin.Context)
	HandleGetPhotos(c *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
}

func NewPhotoController(photoService services.PhotoService) PhotoController {
	return &photoController{photoService}
}

func (p photoController) HandleCreatePhoto(c *gin.Context) {
	claim := c.MustGet("claim").(jwt.MapClaims)
	payload := c.MustGet("payload").(dto.PhotoRequest)

	uid := int(claim["id"].(float64))
	payload.UserID = uint(uid)

	response, err := p.photoService.CreatePhoto(&payload)
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
		Message: "Photo created successfully",
		Data:    response,
	})
}

func (p photoController) HandleUpdatePhoto(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p photoController) HandleDeletePhoto(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p photoController) HandleGetPhotoByID(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p photoController) HandleGetPhotos(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
