package main

import (
	"log"

	"github.com/hambarthegoat/PWEB-FP2-be/config"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/handler"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := config.ConnectSupabase(cfg)

	//REPOS
	userRepo := repository.NewUserRepository(db)
	communityRepo := repository.NewCommunityRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	voteRepo := repository.NewVoteRepository(db)

	//SERVICES
	authSvc := service.NewAuthService(cfg, userRepo)
	userSvc := service.NewUserService(userRepo)
	communitySvc := service.NewCommunityService(communityRepo)
	postSvc := service.NewPostService(postRepo, voteRepo)
	commentSvc := service.NewCommentRepo(commentRepo, voteRepo)
	voteSvc := service.NewVoteRepo(voteRepo)

	//HANDLERS
	r := gin.Default()
	r.Use(middleware.CORS())

	//PUBLIC 
	handler.SetupAuthRoutes(r, authSvc, cfg)
	handler.SetupUserRoutes(r, userSvc)
	handler.SetupCommunityRoutes(r, communitySvc)
	handler.SetupPostRoutes(r, postSvc)
	handler.SetupCommentRoutes(r, commentSvc)
	handler.SetupVoteRoutes(r, voteSvc)

	log.Printf("Server running on :%s", cfg.Port)
	r.Run(":" + cfg.Port)
}
