package dto

// CommentRequest is the request payload for the comment endpoint
type CommentRequest struct {
	Message string `json:"message" validate:"required" example:"Nice photo!"`
	UserID  int    `json:"user_id" swaggerignore:"true"`
	PhotoID int    `json:"photo_id" swaggerignore:"true"`
}

type CommentResponse struct {
	ID        int    `json:"id" example:"1"`
	Message   string `json:"message" example:"Nice photo!"`
	UserID    int    `json:"user_id" example:"1"`
	PhotoID   int    `json:"photo_id" example:"1"`
	CreatedAt string `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}
