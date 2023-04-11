package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my-gram/domain/dto"
	"my-gram/services"
	"net/http"
	"strconv"
)

type CommentController interface {
	HandleCreateComment(c *gin.Context)
	HandleUpdateComment(c *gin.Context)
	HandleDeleteComment(c *gin.Context)
	HandleGetCommentByID(c *gin.Context)
	HandleGetComments(c *gin.Context)
}

type commentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) CommentController {
	return &commentController{commentService}
}

func (ctr commentController) HandleCreateComment(c *gin.Context) {
	claim := c.MustGet("claim").(jwt.MapClaims)
	payload := c.MustGet("payload").(dto.CommentRequest)

	uid := int(claim["id"].(float64))
	pid, _ := strconv.Atoi(c.Param("photoId"))
	payload.UserID = uid
	payload.PhotoID = pid

	response, err := ctr.commentService.CreateComment(&payload)
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
		Message: "Comment created successfully",
		Data:    response,
	})
}

func (ctr commentController) HandleUpdateComment(c *gin.Context) {
	payload := c.MustGet("payload").(dto.CommentRequest)
	cid, _ := strconv.Atoi(c.Param("commentId"))

	response, err := ctr.commentService.UpdateComment(cid, &payload)
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
		Message: "Comment updated successfully",
		Data:    response,
	})
}

func (ctr commentController) HandleDeleteComment(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("commentId"))

	err := ctr.commentService.DeleteComment(cid)
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
		Message: "Comment deleted successfully",
	})
}

func (ctr commentController) HandleGetCommentByID(c *gin.Context) {
	cid, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid comment id. Comment id must be an integer",
		})
		return
	}

	response, errs := ctr.commentService.GetCommentByID(cid)
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
		Message: "Comment retrieved successfully",
		Data:    response,
	})
}

func (ctr commentController) HandleGetComments(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("photoId"))
	response, err := ctr.commentService.GetComments(pid)
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
		Message: "Comments retrieved successfully",
		Data:    response,
	})
}
