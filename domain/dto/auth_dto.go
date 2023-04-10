package dto

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
