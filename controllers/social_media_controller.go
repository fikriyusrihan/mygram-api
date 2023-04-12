package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/services"
	"net/http"
	"strconv"
)

type SocialMediaController interface {
	HandleCreateSocialMedia(c *gin.Context)
	HandleUpdateSocialMedia(c *gin.Context)
	HandleDeleteSocialMedia(c *gin.Context)
	HandleGetSocialMediaByID(c *gin.Context)
	HandleGetSocialMediaByUserID(c *gin.Context)
}

type socialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(socialMediaService services.SocialMediaService) SocialMediaController {
	return &socialMediaController{socialMediaService}
}

// HandleCreateSocialMedia Create social media handler
// @Summary Create Social Media
// @Description Create new social media with name and url. User must be authenticated before using this endpoint.
// @Tags Social Media
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param UserID path int true "User ID"
// @Param Payload body dto.SocialMediaRequest true "Social Media Request Payload"
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.ApiResponse{data=dto.SocialMediaResponse}
// @Router /users/{userId}/social_media [post]
func (ctr socialMediaController) HandleCreateSocialMedia(c *gin.Context) {
	claim := c.MustGet("claim").(jwt.MapClaims)
	payload := c.MustGet("payload").(dto.SocialMediaRequest)

	uid := int(claim["id"].(float64))

	if payload.UserID != uid {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.ApiResponse{
			Code:    http.StatusForbidden,
			Status:  "FORBIDDEN",
			Message: "You are not authorized to create social media for other user",
		})
		return
	}

	response, err := ctr.socialMediaService.CreateSocialMedia(&payload)
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
		Message: "Social media created successfully",
		Data:    response,
	})
}

// HandleUpdateSocialMedia Update social media handler
// @Summary Update Social Media
// @Description Update social media with name and url. User must be authenticated before using this endpoint.
// @Tags Social Media
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param UserID path int true "User ID"
// @Param SocialMediaID path int true "Social Media ID"
// @Param Payload body dto.SocialMediaRequest true "Social Media Request Payload"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.SocialMediaResponse}
// @Router /users/{userId}/social_media/{socialMediaId} [put]
func (ctr socialMediaController) HandleUpdateSocialMedia(c *gin.Context) {
	payload := c.MustGet("payload").(dto.SocialMediaRequest)

	smid, _ := strconv.Atoi(c.Param("socialMediaId"))
	response, err := ctr.socialMediaService.UpdateSocialMedia(smid, &payload)
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
		Message: "Social media updated successfully",
		Data:    response,
	})
}

// HandleDeleteSocialMedia Delete social media handler
// @Summary Delete Social Media
// @Description Delete social media. Only the owner of the user can delete the social media. User must be authenticated before using this endpoint.
// @Tags Social Media
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param UserID path int true "User ID"
// @Param SocialMediaID path int true "Social Media ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse
// @Router /users/{userId}/social_media/{socialMediaId} [delete]
func (ctr socialMediaController) HandleDeleteSocialMedia(c *gin.Context) {
	smid, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := ctr.socialMediaService.DeleteSocialMedia(smid)
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
		Message: "Social media deleted successfully",
	})
}

// HandleGetSocialMediaByID Get social media by user id handler
// @Summary Get Social Media By User ID
// @Description Get social media by user id. User must be authenticated before using this endpoint.
// @Tags Social Media
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param UserID path int true "User ID"
// @Param SocialMediaID path int true "Social Media ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.SocialMediaResponse}
// @Router /users/{userId}/social_media/{socialMediaId} [get]
func (ctr socialMediaController) HandleGetSocialMediaByID(c *gin.Context) {
	smid, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid social media Id. Social media Id must be an integer",
		})
		return
	}

	response, errs := ctr.socialMediaService.GetSocialMediaByID(smid)
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
		Message: "Social media retrieved successfully",
		Data:    response,
	})
}

// HandleGetSocialMediaByUserID Get social media by user id handler
// @Summary Get Social Media By User ID
// @Description Get social media by user id. User must be authenticated before using this endpoint.
// @Tags Social Media
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param UserID path int true "User ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=[]dto.SocialMediaResponse}
// @Router /users/{userId}/social_media [get]
func (ctr socialMediaController) HandleGetSocialMediaByUserID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid user Id. User Id must be an integer",
		})
		return
	}

	response, errs := ctr.socialMediaService.GetSocialMediaByUserID(uid)
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
		Message: "Social media retrieved successfully",
		Data:    response,
	})
}
