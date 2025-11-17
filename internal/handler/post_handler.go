package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"
)

func CreatePostHandler(svc service.PostService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func ListCommunityPostsHandler(svc service.PostService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func UpdatePostHandler(svc service.PostService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }
func DeletePostHandler(svc service.PostService) gin.HandlerFunc { return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) } }

func SetupPostRoutes(r *gin.Engine, svc service.PostService) {
	g := r.Group("/communities/:cid/posts").Use(middleware.Auth())
	g.POST("", CreatePostHandler(svc))
	g.GET("", ListCommunityPostsHandler(svc))

	p := r.Group("/posts").Use(middleware.Auth())
	p.PATCH("/:id", UpdatePostHandler(svc))
	p.DELETE("/:id", DeletePostHandler(svc))
}