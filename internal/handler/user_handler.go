package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"
)

func GetUserHandler(svc service.UserService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func UpdateUserHandler(svc service.UserService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func DeleteUserHandler(svc service.UserService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }

func SetupUserRoutes(r *gin.Engine, svc service.UserService) {
	g := r.Group("/users").Use(middleware.Auth())
	g.GET("/:id", GetUserHandler(svc))
	g.PATCH("/:id", UpdateUserHandler(svc))
	g.DELETE("/:id", DeleteUserHandler(svc))
}