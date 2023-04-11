package repositories

import (
	"gorm.io/gorm"
	"log"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/repositories/repo_interfaces"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repo_interfaces.PhotoRepository {
	return &photoRepository{db}
}

func (p photoRepository) CreatePhoto(photo *entities.Photo) (*entities.Photo, errors.Error) {
	err := p.db.Model(&entities.Photo{}).Create(photo).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return photo, nil
}

func (p photoRepository) UpdatePhoto(id int, photo *entities.Photo) (*entities.Photo, errors.Error) {
	var updatedPhoto *entities.Photo
	err := p.db.Model(&entities.Photo{}).
		Where("id = ?", id).
		Updates(&entities.Photo{
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
		}).
		First(&updatedPhoto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The photo you are trying to update does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return updatedPhoto, nil
}

func (p photoRepository) DeletePhoto(id int) errors.Error {
	var photo entities.Photo
	err := p.db.Model(&entities.Photo{}).Where("id = ?", id).First(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The photo you are trying to delete does not exist")
			return errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	err = p.db.Model(&entities.Photo{}).Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	return nil
}

func (p photoRepository) GetPhotoByID(id int) (*entities.Photo, errors.Error) {
	var photo *entities.Photo
	err := p.db.Model(&entities.Photo{}).Where("id = ?", id).First(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The photo you are trying to get does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return photo, nil
}

func (p photoRepository) GetPhotos() ([]entities.Photo, errors.Error) {
	var photos []entities.Photo
	err := p.db.Model(&entities.Photo{}).Find(&photos).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	if len(photos) == 0 {
		photos = []entities.Photo{}
	}
	return photos, nil
}
