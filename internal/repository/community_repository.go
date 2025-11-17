package repository

import (
	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"gorm.io/gorm"
)

type CommunityRepository interface {
	Create(c *model.Community) error
	FindByID(id string) (*model.Community, error)
	Update(c *model.Community) error
	Delete(id string) error
}

type communityRepository struct{ db *gorm.DB }

func NewCommunityRepository(db *gorm.DB) CommunityRepository { return &communityRepository{db} }

func (r *communityRepository) Create(c *model.Community) error { return r.db.Create(c).Error }
func (r *communityRepository) FindByID(id string) (*model.Community, error) { var c model.Community; err := r.db.Preload("Creator").First(&c, "id = ?", id).Error; return &c, err }
func (r *communityRepository) Update(c *model.Community) error { return r.db.Save(c).Error }
func (r *communityRepository) Delete(id string) error { return r.db.Delete(&model.Community{}, "id = ?", id).Error }