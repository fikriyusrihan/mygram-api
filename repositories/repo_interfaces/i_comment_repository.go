package repo_interfaces

import (
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
)

type CommentRepository interface {
	CreateComment(comment *entities.Comment) (*entities.Comment, errors.Error)
	UpdateComment(id int, comment *entities.Comment) (*entities.Comment, errors.Error)
	DeleteComment(id int) errors.Error
	GetCommentByID(id int) (*entities.Comment, errors.Error)
	GetComments(pid int) ([]entities.Comment, errors.Error)
}
