package belajargolanggin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// mengambil token
		token := c.GetHeader("Authorization")

		if token != "Bearer secret" {
			c.JSON(401, gin.H{
				"message": "token invalid",
			})
			c.Abort()
			return
		}

		//token valid, lanjut boyyy
		c.Next()
	}
}

func TestMiddlewareAuth(t *testing.T) {

	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"token": "secret",
		})
	})

	protected := r.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user":  "humam",
				"email": "admin@admin.com",
			})
		})

		protected.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": []string{"humam", "user2", "user3"},
			})
		})
	}

	r.Run(":8888")
}
