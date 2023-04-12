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

// HandleCreateComment Create comment handler
// @Summary Create Comment
// @Description Create new comment with comment text. User must be authenticated before using this endpoint.
// @Tags Comments
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Param Payload body dto.CommentRequest true "Comment Request Payload"
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.ApiResponse{data=dto.CommentResponse}
// @Router /photos/{photoId}/comments [post]
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

// HandleUpdateComment Update comment handler
// @Summary Update Comment
// @Description Update comment with comment text. User must be authenticated before using this endpoint.
// @Tags Comments
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Param CommentId path int true "Comment ID"
// @Param Payload body dto.CommentRequest true "Comment Request Payload"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.CommentResponse}
// @Router /photos/{photoId}/comments/{commentId} [put]
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

// HandleDeleteComment Delete comment handler
// @Summary Delete Comment
// @Description Delete comment. Only comment owner can delete the comment. User must be authenticated before using this endpoint.
// @Tags Comments
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Param CommentId path int true "Comment ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse
// @Router /photos/{photoId}/comments/{commentId} [delete]
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

// HandleGetCommentByID Get comment by id handler
// @Summary Get Comment By ID
// @Description Get comment by id. User must be authenticated before using this endpoint.
// @Tags Comments
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Param CommentId path int true "Comment ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=dto.CommentResponse}
// @Router /photos/{photoId}/comments/{commentId} [get]
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

// HandleGetComments Get comments handler
// @Summary Get Comments
// @Description Get comments by photo id. User must be authenticated before using this endpoint.
// @Tags Comments
// @Security Bearer
// @Param Authorization header string true "Authentication Bearer Token"
// @Param PhotoId path int true "Photo ID"
// @Produce  json
// @Success 200 {object} dto.ApiResponse{data=[]dto.CommentResponse}
// @Router /photos/{photoId}/comments [get]
func (ctr commentController) HandleGetComments(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid photo id. Photo id must be an integer",
		})
		return
	}

	response, errs := ctr.commentService.GetComments(pid)
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
		Message: "Comments retrieved successfully",
		Data:    response,
	})
}
