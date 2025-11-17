package service

import (
	"errors"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type CommunityService interface {
	Create(req dto.CreateCommunityRequest, userID string) (*dto.CommunityResponse, error)
	GetByID(id string) (*dto.CommunityResponse, error)
	Update(id string, req dto.UpdateCommunityRequest, userID string) (*dto.CommunityResponse, error)
	Delete(id string, userID string) error
}

type communityService struct{ repo repository.CommunityRepository }

func NewCommunityService(r repository.CommunityRepository) CommunityService { return &communityService{repo: r} }

// Only creator can update/delete

func (s *communityService) Create(req dto.CreateCommunityRequest, userID string) (*dto.CommunityResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *communityService) GetByID(id string) (*dto.CommunityResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *communityService) Update(id string, req dto.UpdateCommunityRequest, userID string) (*dto.CommunityResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *communityService) Delete(id string, userID string) error {
	return errors.New("not implemented")
}