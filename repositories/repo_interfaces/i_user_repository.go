package repo_interfaces

import (
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
)

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, errors.Error)
	GetUserByUsername(username string) (*entities.User, errors.Error)
	GetUserByEmail(email string) (*entities.User, errors.Error)
}
