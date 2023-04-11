package dto

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
	UserID  int    `json:"user_id"`
	PhotoID int    `json:"photo_id"`
}

type CommentResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	UserID    int    `json:"user_id"`
	PhotoID   int    `json:"photo_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
