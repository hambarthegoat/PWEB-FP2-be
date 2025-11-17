package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"
)

func VotePostHandler(svc service.VoteService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) }
}

func VoteCommentHandler(svc service.VoteService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) }
}

func SetupVoteRoutes(r *gin.Engine, svc service.VoteService) {
	g := r.Group("").Use(middleware.Auth())
	g.POST("/posts/:id/votes", VotePostHandler(svc))
	g.POST("/comments/:id/votes", VoteCommentHandler(svc))
}