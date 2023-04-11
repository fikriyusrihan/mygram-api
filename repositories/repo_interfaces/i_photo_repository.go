package repo_interfaces

import (
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
)

type PhotoRepository interface {
	CreatePhoto(photo *entities.Photo) (*entities.Photo, errors.Error)
	UpdatePhoto(id int, photo *entities.Photo) (*entities.Photo, errors.Error)
	DeletePhoto(id int) errors.Error
	GetPhotoByID(id int) (*entities.Photo, errors.Error)
	GetPhotos() ([]entities.Photo, errors.Error)
}
