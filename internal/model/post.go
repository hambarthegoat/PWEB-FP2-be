package model

import "time"

type Post struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string    `gorm:"size:300;not null"`
	Content     string    `gorm:"type:text"`
	AuthorID    string    `gorm:"type:uuid;index"`
	Author      User      `gorm:"foreignKey:AuthorID"`
	CommunityID string    `gorm:"type:uuid;index"`
	Community   Community `gorm:"foreignKey:CommunityID"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Score       int       `gorm:"default:0"` // denormalized
}
