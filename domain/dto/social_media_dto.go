package dto

type SocialMediaRequest struct {
	Name   string `json:"name" validate:"required"`
	URL    string `json:"social_media_url" validate:"required"`
	UserID int    `json:"user_id"`
}

type SocialMediaResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"social_media_url"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
