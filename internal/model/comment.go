package model

import "time"

type Comment struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Content   string    `gorm:"type:text;not null"`
	AuthorID  string    `gorm:"type:uuid;index"`
	Author    User      `gorm:"foreignKey:AuthorID"`
	PostID    string    `gorm:"type:uuid;index"`
	Post      Post      `gorm:"foreignKey:PostID"`
	ParentID  *string   `gorm:"type:uuid;index"` // NULL = root
	Parent    *Comment  `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Score     int       `gorm:"default:0"` // denormalized
}
