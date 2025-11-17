package service

import (
	"errors"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type PostService interface {
	Create(req dto.CreatePostRequest, userID, communityID string) (*dto.PostResponse, error)
	GetCommunityPosts(communityID string, limit, offset int) ([]dto.PostListItem, error)
	Update(id string, req dto.UpdatePostRequest, userID string) (*dto.PostResponse, error)
	Delete(id string, userID string) error
}

type postService struct {
	postRepo repository.PostRepository
	voteRepo repository.VoteRepository
}

func NewPostService(pr repository.PostRepository, vr repository.VoteRepository) PostService {
	return &postService{postRepo: pr, voteRepo: vr}
}

// Join vote count → denormalized Score

func (s *postService) Create(req dto.CreatePostRequest, userID, communityID string) (*dto.PostResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *postService) GetCommunityPosts(communityID string, limit, offset int) ([]dto.PostListItem, error) {
	return nil, errors.New("not implemented")
}

func (s *postService) Update(id string, req dto.UpdatePostRequest, userID string) (*dto.PostResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *postService) Delete(id string, userID string) error {
	return errors.New("not implemented")
}