package entities

import "my-gram/domain/dto"

type Photo struct {
	GormModel
	Title    string `gorm:"notNull"`
	Caption  string
	PhotoURL string `gorm:"notNull"`
	UserID   uint   `gorm:"notNull;index"`
}

func (p *Photo) ToResponse() dto.PhotoResponse {
	return dto.PhotoResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func (p *Photo) FromRequest(request *dto.PhotoRequest) {
	p.Title = request.Title
	p.Caption = request.Caption
	p.PhotoURL = request.PhotoURL
}
