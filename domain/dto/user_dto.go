package dto

type UserRequest struct {
	Username string `json:"username" validate:"required" example:"fikriyusrihan"`
	Password string `json:"password" validate:"required,min=6" example:"supersecret" minLength:"6"`
	Email    string `json:"email" validate:"required,email" example:"fikriyusrihan@gmail.com"`
	Age      uint   `json:"age" validate:"required,min=8" example:"24" minimum:"8"`
}

type UserResponse struct {
	ID        int    `json:"id" example:"1"`
	Username  string `json:"username" example:"fikriyusrihan"`
	Email     string `json:"email" example:"fikriyusrihan@gmail.com"`
	Age       uint   `json:"age" example:"24"`
	CreatedAt string `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}
