package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hambarthegoat/PWEB-FP2-be/config"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"
)

func RegisterHandler(svc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
		res, err := svc.Register(req.Email, req.Password)
		if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
		c.JSON(http.StatusCreated, res)
	}
}

func LoginHandler(svc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
		res, err := svc.Login(req.Email, req.Password)
		if err != nil { c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}); return }
		c.JSON(http.StatusOK, res)
	}
}

func SetupAuthRoutes(r *gin.Engine, svc service.AuthService, cfg *config.Config) {
	auth := r.Group("/auth")
	auth.POST("/register", RegisterHandler(svc))
	auth.POST("/login", LoginHandler(svc))
	// OAuth endpoints intentionally left unimplemented for now
	auth.POST("/logout", middleware.Auth(), func(c *gin.Context) { c.Status(http.StatusNoContent) })
}