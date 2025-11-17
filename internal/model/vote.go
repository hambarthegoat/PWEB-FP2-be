package model

import "time"

type VoteTargetType string

const (
	VoteTargetPost    VoteTargetType = "post"
	VoteTargetComment VoteTargetType = "comment"
)

type Vote struct {
	UserID     string         `gorm:"type:uuid;primaryKey"`
	TargetID   string         `gorm:"type:uuid;primaryKey"`
	TargetType VoteTargetType `gorm:"type:text;primaryKey"`
	Value      int8           `gorm:"not null"` // +1 or -1
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
}
