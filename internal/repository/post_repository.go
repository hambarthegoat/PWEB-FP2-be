package repository

import (
	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(p *model.Post) error
	FindByID(id string) (*model.Post, error)
	FindByCommunityID(cid string, limit, offset int) ([]model.Post, error)
	Update(p *model.Post) error
	Delete(id string) error
	UpdateScore(id string, score int) error
}

type postRepository struct{ db *gorm.DB }

func NewPostRepository(db *gorm.DB) PostRepository { return &postRepository{db} }

func (r *postRepository) Create(p *model.Post) error { return r.db.Create(p).Error }
func (r *postRepository) FindByID(id string) (*model.Post, error) { var p model.Post; err := r.db.Preload("Author").Preload("Community").First(&p, "id = ?", id).Error; return &p, err }
func (r *postRepository) FindByCommunityID(cid string, limit, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Preload("Author").Where("community_id = ?", cid).Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}
func (r *postRepository) Update(p *model.Post) error { return r.db.Save(p).Error }
func (r *postRepository) Delete(id string) error { return r.db.Delete(&model.Post{}, "id = ?", id).Error }
func (r *postRepository) UpdateScore(id string, score int) error { return r.db.Model(&model.Post{}).Where("id = ?", id).Update("score", score).Error }