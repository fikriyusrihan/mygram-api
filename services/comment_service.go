package services

import (
	"my-gram/domain/dto"
	"my-gram/domain/entities"
	"my-gram/pkg/errors"
	"my-gram/repositories/repo_interfaces"
)

type CommentService interface {
	CreateComment(payload *dto.CommentRequest) (*dto.CommentResponse, errors.Error)
	UpdateComment(id int, payload *dto.CommentRequest) (*dto.CommentResponse, errors.Error)
	DeleteComment(id int) errors.Error
	GetCommentByID(id int) (*dto.CommentResponse, errors.Error)
	GetComments(pid int) ([]dto.CommentResponse, errors.Error)
}

type commentService struct {
	commentRepository repo_interfaces.CommentRepository
	photoRepository   repo_interfaces.PhotoRepository
}

func NewCommentService(
	commentRepository repo_interfaces.CommentRepository,
	photoRepository repo_interfaces.PhotoRepository,
) CommentService {
	return &commentService{
		commentRepository,
		photoRepository,
	}
}

func (c commentService) CreateComment(payload *dto.CommentRequest) (*dto.CommentResponse, errors.Error) {
	var comment entities.Comment
	comment.FromRequest(payload)

	createdComment, err := c.commentRepository.CreateComment(&comment)
	if err != nil {
		return nil, err
	}

	return createdComment.ToResponse(), nil
}

func (c commentService) UpdateComment(id int, payload *dto.CommentRequest) (*dto.CommentResponse, errors.Error) {
	var comment entities.Comment
	comment.FromRequest(payload)

	updatedComment, err := c.commentRepository.UpdateComment(id, &comment)
	if err != nil {
		return nil, err
	}

	return updatedComment.ToResponse(), nil
}

func (c commentService) DeleteComment(id int) errors.Error {
	err := c.commentRepository.DeleteComment(id)
	if err != nil {
		return err
	}

	return nil
}

func (c commentService) GetCommentByID(id int) (*dto.CommentResponse, errors.Error) {
	comment, err := c.commentRepository.GetCommentByID(id)
	if err != nil {
		return nil, err
	}

	return comment.ToResponse(), nil
}

func (c commentService) GetComments(pid int) ([]dto.CommentResponse, errors.Error) {
	_, err := c.photoRepository.GetPhotoByID(pid)
	if err != nil {
		return nil, err
	}

	comments, err := c.commentRepository.GetComments(pid)
	if err != nil {
		return nil, err
	}

	var responses []dto.CommentResponse
	for _, comment := range comments {
		responses = append(responses, *comment.ToResponse())
	}

	return responses, nil
}
