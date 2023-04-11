package repositories

import (
	"gorm.io/gorm"
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
		panic(err)
	}

	return photo, nil
}

func (p photoRepository) UpdatePhoto(id int, photo *entities.Photo) (*entities.Photo, errors.Error) {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) DeletePhoto(id int) errors.Error {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) GetPhotoByID(id int) (*entities.Photo, errors.Error) {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) GetPhotos() ([]entities.Photo, errors.Error) {
	//TODO implement me
	panic("implement me")
}
