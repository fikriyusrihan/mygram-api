package repo_interfaces

import (
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *entities.SocialMedia) (*entities.SocialMedia, errors.Error)
	UpdateSocialMedia(id int, socialMedia *entities.SocialMedia) (*entities.SocialMedia, errors.Error)
	DeleteSocialMedia(id int) errors.Error
	GetSocialMediaByID(id int) (*entities.SocialMedia, errors.Error)
	GetSocialMediaByUserID(uid int) ([]*entities.SocialMedia, errors.Error)
}
