package services

import (
	"my-gram/domain/dto"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/pkg/helpers"
	"my-gram/repositories/repo_interfaces"
)

type AuthService interface {
	Login(payload *dto.AuthRequest) (*dto.AuthResponse, errors.Error)
	Register(payload *dto.UserRequest) (*dto.UserResponse, errors.Error)
}

type authService struct {
	userRepository repo_interfaces.UserRepository
}

func NewAuthService(userRepository repo_interfaces.UserRepository) AuthService {
	return &authService{userRepository}
}

func (a authService) Login(payload *dto.AuthRequest) (*dto.AuthResponse, errors.Error) {
	user, errs := a.userRepository.GetUserByUsername(payload.Username)
	if errs != nil {
		if errs.Code() == 404 {
			userByEmail, errsByEmail := a.userRepository.GetUserByEmail(payload.Username)
			if errsByEmail != nil {
				return nil, errsByEmail
			}

			user = userByEmail
		} else {
			return nil, errs
		}
	}

	isValidPassword := helpers.ValidatePassword(user.Password, payload.Password)
	if !isValidPassword {
		errs = errors.NewUnauthenticatedError("Invalid username or password. Please check your username and password and try again")
		return nil, errs
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		errs = errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	response := &dto.AuthResponse{
		AccessToken: token,
	}

	return response, nil
}

func (a authService) Register(payload *dto.UserRequest) (*dto.UserResponse, errors.Error) {
	var user entities.User
	user.FromRequest(payload)

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	user.Password = hashedPassword
	result, errs := a.userRepository.CreateUser(&user)
	if errs != nil {
		return nil, errs
	}

	response := result.ToResponse()
	return response, nil
}
