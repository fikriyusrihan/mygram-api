package repositories

import (
	"gorm.io/gorm"
	"log"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/repositories/repo_interfaces"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) repo_interfaces.SocialMediaRepository {
	return &socialMediaRepository{db}
}

func (s socialMediaRepository) CreateSocialMedia(socialMedia *entities.SocialMedia) (*entities.SocialMedia, errors.Error) {
	err := s.db.Model(&entities.SocialMedia{}).Create(socialMedia).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return socialMedia, nil
}

func (s socialMediaRepository) UpdateSocialMedia(id int, socialMedia *entities.SocialMedia) (*entities.SocialMedia, errors.Error) {
	var updatedSocialMedia *entities.SocialMedia
	err := s.db.Model(&entities.SocialMedia{}).
		Where("id = ?", id).
		Updates(&entities.SocialMedia{
			Name: socialMedia.Name,
			URL:  socialMedia.URL,
		}).
		First(&updatedSocialMedia).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The social media you are trying to update does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return updatedSocialMedia, nil
}

func (s socialMediaRepository) DeleteSocialMedia(id int) errors.Error {
	var socialMedia entities.SocialMedia
	err := s.db.Model(&entities.SocialMedia{}).Where("id = ?", id).First(&socialMedia).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The social media you are trying to delete does not exist")
			return errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	err = s.db.Model(&entities.SocialMedia{}).Where("id = ?", id).Delete(&socialMedia).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	return nil
}

func (s socialMediaRepository) GetSocialMediaByID(id int) (*entities.SocialMedia, errors.Error) {
	var socialMedia entities.SocialMedia
	err := s.db.Model(&entities.SocialMedia{}).Where("id = ?", id).First(&socialMedia).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The social media you are trying to retrieve does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return &socialMedia, nil
}

func (s socialMediaRepository) GetSocialMediaByUserID(uid int) ([]*entities.SocialMedia, errors.Error) {
	var socialMedias []*entities.SocialMedia
	err := s.db.Model(&entities.SocialMedia{}).Where("user_id = ?", uid).Find(&socialMedias).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	if len(socialMedias) == 0 {
		socialMedias = []*entities.SocialMedia{}
	}

	return socialMedias, nil
}
