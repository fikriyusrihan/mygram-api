package dto

// AuthRequest is the request payload for the login endpoint
type AuthRequest struct {
	Email    string `json:"email" validate:"required" example:"fikriyusrihan@gmail.com"`
	Password string `json:"password" validate:"required,min=6" example:"supersecret"`
}

// AuthResponse is the response payload for the login endpoint
type AuthResponse struct {
	UserID      int    `json:"user_id" example:"1"`
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIn0.ri7_-S3RIefxm6JxzsJSWVSyvSTOIivZgcuVDqaR3fQ"`
}
