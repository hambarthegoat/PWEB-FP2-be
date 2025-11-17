package service

import (
	"errors"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type UserService interface {
	GetByID(id string) (*dto.UserResponse, error)
	Update(id string, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(id string) error
}

type userService struct{ userRepo repository.UserRepository }

func NewUserService(r repository.UserRepository) UserService { return &userService{userRepo: r} }

// Soft-delete or hard (cascade vote/comment/post)

func (s *userService) GetByID(id string) (*dto.UserResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *userService) Update(id string, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *userService) Delete(id string) error {
	return errors.New("not implemented")
}