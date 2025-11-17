package repository

import (
	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByOAuth(provider, sub string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository { return &userRepository{db} }

func (r *userRepository) Create(u *model.User) error { return r.db.Create(u).Error }
func (r *userRepository) FindByID(id string) (*model.User, error) { var u model.User; err := r.db.First(&u, "id = ?", id).Error; return &u, err }
func (r *userRepository) FindByEmail(email string) (*model.User, error) { var u model.User; err := r.db.First(&u, "email = ?", email).Error; return &u, err }
func (r *userRepository) FindByOAuth(provider, sub string) (*model.User, error) { var u model.User; err := r.db.First(&u, "provider = ? AND sub = ?", provider, sub).Error; return &u, err }
func (r *userRepository) Update(u *model.User) error { return r.db.Save(u).Error }
func (r *userRepository) Delete(id string) error { return r.db.Delete(&model.User{}, "id = ?", id).Error }