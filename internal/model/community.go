package model

import "time"

type Community struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"uniqueIndex;size:100;not null"`
	Description string    `gorm:"size:500"`
	CreatorID   string    `gorm:"type:uuid;index"`
	Creator     User      `gorm:"foreignKey:CreatorID"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
