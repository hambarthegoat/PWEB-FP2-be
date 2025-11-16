package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectSupabase(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.SupabaseDBURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect supabase: " + err.Error())
	}
	return db
}