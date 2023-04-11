package repositories

import (
	"gorm.io/gorm"
	"log"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/pkg/helpers"
	"my-gram/repositories/repo_interfaces"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repo_interfaces.CommentRepository {
	return &commentRepository{db}
}

func (c commentRepository) CreateComment(comment *entities.Comment) (*entities.Comment, errors.Error) {
	err := c.db.Model(&entities.Comment{}).Create(comment).Error
	if err != nil {
		if helpers.IsForeignKeyViolation(err) {
			errs := errors.NewBadRequestError("The photo you are trying to comment on does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return comment, nil
}

func (c commentRepository) UpdateComment(id int, comment *entities.Comment) (*entities.Comment, errors.Error) {
	var updatedComment entities.Comment

	err := c.db.Model(&entities.Comment{}).
		Where("id = ?", id).
		Updates(&entities.Comment{
			Message: comment.Message,
		}).
		First(&updatedComment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The comment you are trying to update does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return &updatedComment, nil
}

func (c commentRepository) DeleteComment(id int) errors.Error {
	var comment entities.Comment
	err := c.db.Model(&entities.Comment{}).Where("id = ?", id).First(&comment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The comment you are trying to delete does not exist")
			return errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	err = c.db.Model(&entities.Comment{}).Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	return nil
}

func (c commentRepository) GetCommentByID(id int) (*entities.Comment, errors.Error) {
	var comment entities.Comment
	err := c.db.Model(&entities.Comment{}).Where("id = ?", id).First(&comment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The comment you are trying to get does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return &comment, nil
}

func (c commentRepository) GetComments(pid int) ([]entities.Comment, errors.Error) {
	var comments []entities.Comment
	err := c.db.Model(&entities.Comment{}).Where("photo_id = ?", pid).Find(&comments).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return comments, nil
}
