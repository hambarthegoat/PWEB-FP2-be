package middleware

import "github.com/gin-gonic/gin"

func SupabaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}

		// TODO: verify Supabase JWT
		
		c.Next()
	}
}
