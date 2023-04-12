package entities

import "my-gram/domain/dto"

type SocialMedia struct {
	GormModel
	Name   string `gorm:"notNull"`
	URL    string `gorm:"notNull"`
	UserID int    `gorm:"notNull;index"`
}

func (s *SocialMedia) ToResponse() dto.SocialMediaResponse {
	return dto.SocialMediaResponse{
		ID:        s.ID,
		Name:      s.Name,
		URL:       s.URL,
		UserID:    s.UserID,
		CreatedAt: s.CreatedAt.String(),
		UpdatedAt: s.UpdatedAt.String(),
	}
}

func (s *SocialMedia) FromRequest(payload *dto.SocialMediaRequest) {
	s.Name = payload.Name
	s.URL = payload.URL
	s.UserID = payload.UserID
}
