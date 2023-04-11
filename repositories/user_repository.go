package repositories

import (
	"gorm.io/gorm"
	"log"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/repositories/repo_interfaces"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repo_interfaces.UserRepository {
	return &userRepository{db}
}

func (u userRepository) CreateUser(user *entities.User) (*entities.User, errors.Error) {
	err := u.db.Model(&entities.User{}).Create(user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			errs := errors.NewConflictError("User with this email or username already exists. Please try again with a different email or username")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return user, nil
}

func (u userRepository) GetUserByUsername(username string) (*entities.User, errors.Error) {
	var user entities.User

	err := u.db.Model(&entities.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewUnauthenticatedError("Invalid username or password. Please check your username and password and try again")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return &user, nil
}

func (u userRepository) GetUserByEmail(email string) (*entities.User, errors.Error) {
	var user entities.User

	err := u.db.Model(&entities.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewUnauthenticatedError("Invalid email or password. Please check your email and password and try again")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return &user, nil
}
