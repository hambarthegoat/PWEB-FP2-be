package service

import (
	"errors"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type CommentService interface {
	Create(req dto.CreateCommentRequest, userID, postID string) (*dto.CommentResponse, error)
	GetByPostID(postID string, parentID *string) ([]dto.CommentTree, error)
	Update(id string, req dto.UpdateCommentRequest, userID string) (*dto.CommentResponse, error)
	Delete(id string, userID string) error
}

type commentService struct {
	commentRepo repository.CommentRepository
	voteRepo    repository.VoteRepository
}

func NewCommentService(cr repository.CommentRepository, vr repository.VoteRepository) CommentService {
	return &commentService{commentRepo: cr, voteRepo: vr}
}

// Build tree via recursive query or in-memory

func (s *commentService) Create(req dto.CreateCommentRequest, userID, postID string) (*dto.CommentResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *commentService) GetByPostID(postID string, parentID *string) ([]dto.CommentTree, error) {
	return nil, errors.New("not implemented")
}

func (s *commentService) Update(id string, req dto.UpdateCommentRequest, userID string) (*dto.CommentResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *commentService) Delete(id string, userID string) error {
	return errors.New("not implemented")
}