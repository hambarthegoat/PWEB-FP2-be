package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"
)

func CreateCommunityHandler(svc service.CommunityService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func GetCommunityHandler(svc service.CommunityService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func UpdateCommunityHandler(svc service.CommunityService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func DeleteCommunityHandler(svc service.CommunityService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }

func SetupCommunityRoutes(r *gin.Engine, svc service.CommunityService) {
	g := r.Group("/communities").Use(middleware.Auth())
	g.POST("", CreateCommunityHandler(svc))
	g.GET("/:id", GetCommunityHandler(svc))
	g.PATCH("/:id", UpdateCommunityHandler(svc))
	g.DELETE("/:id", DeleteCommunityHandler(svc))
}