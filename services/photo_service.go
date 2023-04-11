package services

import (
	"my-gram/domain/dto"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/repositories/repo_interfaces"
)

type PhotoService interface {
	CreatePhoto(payload *dto.PhotoRequest) (*dto.PhotoResponse, errors.Error)
	UpdatePhoto(id int, payload *dto.PhotoRequest) (*dto.PhotoResponse, errors.Error)
	DeletePhoto(id int) errors.Error
	GetPhotoByID(id int) (*dto.PhotoResponse, errors.Error)
	GetPhotos() ([]dto.PhotoResponse, errors.Error)
}

type photoService struct {
	photoRepository repo_interfaces.PhotoRepository
}

func NewPhotoService(photoRepository repo_interfaces.PhotoRepository) PhotoService {
	return &photoService{photoRepository}
}

func (p photoService) CreatePhoto(payload *dto.PhotoRequest) (*dto.PhotoResponse, errors.Error) {
	var photo entities.Photo
	photo.FromRequest(payload)

	createdPhoto, err := p.photoRepository.CreatePhoto(&photo)
	if err != nil {
		return nil, err
	}

	response := createdPhoto.ToResponse()

	return &response, nil
}

func (p photoService) UpdatePhoto(id int, payload *dto.PhotoRequest) (*dto.PhotoResponse, errors.Error) {
	//TODO implement me
	panic("implement me")
}

func (p photoService) DeletePhoto(id int) errors.Error {
	//TODO implement me
	panic("implement me")
}

func (p photoService) GetPhotoByID(id int) (*dto.PhotoResponse, errors.Error) {
	//TODO implement me
	panic("implement me")
}

func (p photoService) GetPhotos() ([]dto.PhotoResponse, errors.Error) {
	//TODO implement me
	panic("implement me")
}
