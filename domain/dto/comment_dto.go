package dto

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentResponse struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	UserID    uint   `json:"user_id"`
	PhotoID   uint   `json:"photo_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
