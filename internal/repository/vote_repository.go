package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
)

type VoteRepository interface {
	Upsert(vote *model.Vote) error
	GetUserVote(userID, targetID string, targetType model.VoteTargetType) (*model.Vote, error)
	CalculateScore(targetID string, targetType model.VoteTargetType) (int, error)
}
type voteRepository struct{ db *gorm.DB }

func NewVoteRepository(db *gorm.DB) VoteRepository { return &voteRepository{db} }

func (r *voteRepository) Upsert(v *model.Vote) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "target_id"}, {Name: "target_type"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": v.Value, "updated_at": time.Now()}),
	}).Create(v).Error
}

func (r *voteRepository) GetUserVote(userID, targetID string, targetType model.VoteTargetType) (*model.Vote, error) {
	var v model.Vote
	err := r.db.First(&v, "user_id = ? AND target_id = ? AND target_type = ?", userID, targetID, targetType).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { return nil, nil }
	return &v, err
}

func (r *voteRepository) CalculateScore(targetID string, targetType model.VoteTargetType) (int, error) {
	var score int
	err := r.db.Model(&model.Vote{}).
		Select("COALESCE(SUM(value), 0)").
		Where("target_id = ? AND target_type = ?", targetID, targetType).
		Scan(&score).Error
	return score, err
}