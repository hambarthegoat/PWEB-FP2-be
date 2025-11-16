package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL string
	SupabaseAnon string
	SupabaseDBURL string
	JWTSecret string
	Port string
}

func Load() *Config {
	godotenv.Load()
	return &Config {
		SupabaseURL:  os.Getenv("SUPABASE_URL"),
		SupabaseAnon: os.Getenv("SUPABASE_ANON_KEY"),
		SupabaseDBURL: os.Getenv("SUPABASE_DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port: os.Getenv("PORT"),
	}
}

func getEnv(key, fallback string) string{
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}