package services

import (
	"my-gram/domain/dto"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/repositories/repo_interfaces"
)

type SocialMediaService interface {
	CreateSocialMedia(payload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errors.Error)
	UpdateSocialMedia(id int, payload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errors.Error)
	DeleteSocialMedia(id int) errors.Error
	GetSocialMediaByID(id int) (*dto.SocialMediaResponse, errors.Error)
	GetSocialMediaByUserID(uid int) ([]dto.SocialMediaResponse, errors.Error)
}

type socialMediaService struct {
	socialMediaRepository repo_interfaces.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repo_interfaces.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepository}
}

func (s socialMediaService) CreateSocialMedia(payload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errors.Error) {
	var socialMedia entities.SocialMedia
	socialMedia.FromRequest(payload)

	createdSocialMedia, err := s.socialMediaRepository.CreateSocialMedia(&socialMedia)
	if err != nil {
		return nil, err
	}

	response := createdSocialMedia.ToResponse()
	return &response, nil
}

func (s socialMediaService) UpdateSocialMedia(id int, payload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errors.Error) {
	var socialMedia entities.SocialMedia
	socialMedia.FromRequest(payload)

	updatedSocialMedia, err := s.socialMediaRepository.UpdateSocialMedia(id, &socialMedia)
	if err != nil {
		return nil, err
	}

	response := updatedSocialMedia.ToResponse()
	return &response, nil
}

func (s socialMediaService) DeleteSocialMedia(id int) errors.Error {
	err := s.socialMediaRepository.DeleteSocialMedia(id)
	if err != nil {
		return err
	}

	return nil
}

func (s socialMediaService) GetSocialMediaByID(id int) (*dto.SocialMediaResponse, errors.Error) {
	socialMedia, err := s.socialMediaRepository.GetSocialMediaByID(id)
	if err != nil {
		return nil, err
	}

	response := socialMedia.ToResponse()
	return &response, nil
}

func (s socialMediaService) GetSocialMediaByUserID(uid int) ([]dto.SocialMediaResponse, errors.Error) {
	socialMedias, err := s.socialMediaRepository.GetSocialMediaByUserID(uid)
	if err != nil {
		return nil, err
	}

	var response []dto.SocialMediaResponse
	for _, socialMedia := range socialMedias {
		response = append(response, socialMedia.ToResponse())
	}

	return response, nil
}
