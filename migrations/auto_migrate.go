package migrate

import (
	"log"

	"github.com/hambarthegoat/PWEB-FP2-be/config"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
)

func Migrate() {
	cfg := config.Load()
	db := config.ConnectSupabase(cfg)

	err := db.AutoMigrate(
		&model.User{}, 
		&model.Community{},
		&model.Post{},
		&model.Comment{},
		&model.Vote{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration done")
}