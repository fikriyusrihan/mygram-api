package entities

import "my-gram/domain/dto"

type Comment struct {
	GormModel
	Message string `gorm:"notNull"`
	UserID  int    `gorm:"notNull;index"`
	PhotoID int    `gorm:"notNull;index"`
}

func (c *Comment) FromRequest(request *dto.CommentRequest) {
	c.Message = request.Message
	c.UserID = request.UserID
	c.PhotoID = request.PhotoID
}

func (c *Comment) ToResponse() *dto.CommentResponse {
	return &dto.CommentResponse{
		ID:        c.ID,
		Message:   c.Message,
		UserID:    c.UserID,
		PhotoID:   c.PhotoID,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}
