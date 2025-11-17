package model

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string    `gorm:"uniqueIndex;size:255;not null"`
	Password  *string   `gorm:"size:255"`             // NULL for OAuth-only
	Provider  *string   `gorm:"size:50"`              // "google", "github", NULL
	Sub       *string   `gorm:"uniqueIndex;size:255"` // OAuth subject
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
