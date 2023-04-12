package dto

type SocialMediaRequest struct {
	Name   string `json:"name" validate:"required" example:"Instagram"`
	URL    string `json:"social_media_url" validate:"required" example:"https://instagram.com/fikriyusrihan"`
	UserID int    `json:"user_id" swaggerignore:"true"`
}

type SocialMediaResponse struct {
	ID        int    `json:"id" example:"1"`
	Name      string `json:"name" example:"Instagram"`
	URL       string `json:"social_media_url" example:"https://instagram.com/fikriyusrihan"`
	UserID    int    `json:"user_id" example:"1"`
	CreatedAt string `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}
