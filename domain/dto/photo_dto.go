package dto

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
	UserID   int    `json:"user_id"`
}

type PhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type PhotoDetailResponse struct {
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	PhotoURL  string            `json:"photo_url"`
	UserID    int               `json:"user_id"`
	Comments  []CommentResponse `json:"comments"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
