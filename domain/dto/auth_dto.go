package dto

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	UserID      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
}
