package entities

import "my-gram/domain/dto"

type Photo struct {
	GormModel
	Title    string `gorm:"notNull"`
	Caption  string
	PhotoURL string `gorm:"notNull"`
	UserID   int    `gorm:"notNull;index"`
	Comments []Comment
}

func (p *Photo) ToResponse() dto.PhotoResponse {
	return dto.PhotoResponse{
		ID:       p.ID,
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
		UserID:   p.UserID,
	}
}

func (p *Photo) ToDetailResponse() dto.PhotoDetailResponse {
	var comments []dto.CommentResponse
	for _, c := range p.Comments {
		comments = append(comments, *c.ToResponse())
	}

	return dto.PhotoDetailResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		UserID:    p.UserID,
		Comments:  comments,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func (p *Photo) FromRequest(request *dto.PhotoRequest) {
	p.Title = request.Title
	p.Caption = request.Caption
	p.PhotoURL = request.PhotoURL
	p.UserID = request.UserID
}
