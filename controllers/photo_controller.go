package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/services"
	"net/http"
	"strconv"
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

// HandleCreatePhoto Create photo handler
// @Summary Create Photo
// @Description Create new photo with title, caption, and photo url. User must be authenticated before using this endpoint.
// @Tags Photos
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param Payload body dto.PhotoRequest true "Photo Request Payload"
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.ApiResponse{data=dto.PhotoResponse}
// @Router /photos [post]
func (ctr photoController) HandleCreatePhoto(c *gin.Context) {
	claim := c.MustGet("claim").(jwt.MapClaims)
	payload := c.MustGet("payload").(dto.PhotoRequest)

	uid := int(claim["id"].(float64))
	payload.UserID = uid

	response, err := ctr.photoService.CreatePhoto(&payload)
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

// HandleUpdatePhoto Update photo handler
// @Summary Update Photo
// @Description Update photo with title, caption, and photo url. User must be authenticated before using this endpoint.
// @Tags Photos
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Param Payload body dto.PhotoRequest true "Photo Request Payload"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.PhotoResponse}
// @Router /photos/{photoId} [put]
func (ctr photoController) HandleUpdatePhoto(c *gin.Context) {
	claim := c.MustGet("claim").(jwt.MapClaims)
	payload := c.MustGet("payload").(dto.PhotoRequest)

	uid := int(claim["id"].(float64))
	pid, _ := strconv.Atoi(c.Param("photoId"))
	payload.UserID = uid

	response, errs := ctr.photoService.UpdatePhoto(pid, &payload)
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Photo updated successfully",
		Data:    response,
	})
}

// HandleDeletePhoto Delete photo handler
// @Summary Delete Photo
// @Description Delete photo by photo id. Only the owner of the photo can delete it. User must be authenticated before using this endpoint.
// @Tags Photos
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ApiResponse
// @Router /photos/{photoId} [delete]
func (ctr photoController) HandleDeletePhoto(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("photoId"))

	errs := ctr.photoService.DeletePhoto(pid)
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Photo deleted successfully",
	})
}

// HandleGetPhotoByID Get photo by id handler
// @Summary Get Photo By ID
// @Description Get photo by photo id. User must be authenticated before using this endpoint.
// @Tags Photos
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.PhotoDetailResponse}
// @Router /photos/{photoId} [get]
func (ctr photoController) HandleGetPhotoByID(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid photo id. Photo id must be an integer",
		})
		return
	}

	response, errs := ctr.photoService.GetPhotoByID(pid)
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Photo retrieved successfully",
		Data:    response,
	})
}

// HandleGetPhotos Get photos handler
// @Summary Get Photos
// @Description Get all photos. User must be authenticated before using this endpoint.
// @Tags Photos
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=[]dto.PhotoResponse}
// @Router /photos [get]
func (ctr photoController) HandleGetPhotos(c *gin.Context) {
	response, errs := ctr.photoService.GetPhotos()
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Photos retrieved successfully",
		Data:    response,
	})
}
