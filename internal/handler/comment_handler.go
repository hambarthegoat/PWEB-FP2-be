package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/service"
	"github.com/hambarthegoat/PWEB-FP2-be/pkg/middleware"
)

func CreateCommentHandler(svc service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) }
}

func ListCommentsHandler(svc service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) }
}

func UpdateCommentHandler(svc service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) }
}

func DeleteCommentHandler(svc service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"}) }
}

func SetupCommentRoutes(r *gin.Engine, svc service.CommentService) {
	g := r.Group("/posts/:pid/comments").Use(middleware.Auth())
	g.POST("", CreateCommentHandler(svc))
	g.GET("", ListCommentsHandler(svc))

	c := r.Group("/comments").Use(middleware.Auth())
	c.PATCH("/:id", UpdateCommentHandler(svc))
	c.DELETE("/:id", DeleteCommentHandler(svc))
}