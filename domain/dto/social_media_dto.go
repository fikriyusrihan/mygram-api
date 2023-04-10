package dto

type SocialMediaRequest struct {
	Name string `json:"name" validate:"required"`
	URL  string `json:"social_media_url" validate:"required"`
}

type SocialMediaResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"social_media_url"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
