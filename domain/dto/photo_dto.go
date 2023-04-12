package dto

type PhotoRequest struct {
	Title    string `json:"title" validate:"required" example:"Vacation in Bali"`
	Caption  string `json:"caption" example:"I had a great time in Bali!"`
	PhotoURL string `json:"photo_url" validate:"required" example:"https://my-gram.com/photos/1.jpg"`
	UserID   int    `json:"user_id" swaggerignore:"true"`
}

type PhotoResponse struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"Vacation in Bali"`
	Caption  string `json:"caption" example:"I had a great time in Bali!"`
	PhotoURL string `json:"photo_url" example:"https://my-gram.com/photos/1.jpg"`
	UserID   int    `json:"user_id" example:"1"`
}

type PhotoDetailResponse struct {
	ID        int               `json:"id" example:"1"`
	Title     string            `json:"title" example:"Vacation in Bali"`
	Caption   string            `json:"caption" example:"I had a great time in Bali!"`
	PhotoURL  string            `json:"photo_url" example:"https://my-gram.com/photos/1.jpg"`
	UserID    int               `json:"user_id" example:"1"`
	Comments  []CommentResponse `json:"comments"`
	CreatedAt string            `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt string            `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}
