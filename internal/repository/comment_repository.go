package repository

import (
	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(c *model.Comment) error
	FindByPostID(postID string, parentID *string) ([]model.Comment, error)
	FindByID(id string) (*model.Comment, error)
	Update(c *model.Comment) error
	Delete(id string) error
	UpdateScore(id string, score int) error
}

type commentRepository struct{ db *gorm.DB }

func NewCommentRepository(db *gorm.DB) CommentRepository { return &commentRepository{db} }

func (r *commentRepository) Create(c *model.Comment) error { return r.db.Create(c).Error }
func (r *commentRepository) FindByPostID(postID string, parentID *string) ([]model.Comment, error) {
	var comments []model.Comment
	query := r.db.Preload("Author").Where("post_id = ?", postID)
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}
	err := query.Order("created_at ASC").Find(&comments).Error
	return comments, err
}
func (r *commentRepository) FindByID(id string) (*model.Comment, error) { var c model.Comment; err := r.db.Preload("Author").First(&c, "id = ?", id).Error; return &c, err }
func (r *commentRepository) Update(c *model.Comment) error { return r.db.Save(c).Error }
func (r *commentRepository) Delete(id string) error { return r.db.Delete(&model.Comment{}, "id = ?", id).Error }
func (r *commentRepository) UpdateScore(id string, score int) error { return r.db.Model(&model.Comment{}).Where("id = ?", id).Update("score", score).Error }