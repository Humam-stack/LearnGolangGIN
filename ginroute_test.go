package belajargolanggin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRouterGin(t *testing.T) {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get All user",
			"data":    []string{"humam", "user1", "user2"},
		})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Berhasil membuat user baru",
			"data": gin.H{
				"username": "user3",
				"email":    "user3@user3.com",
			},
		})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "user updated! (full)",
			"id":      id,
		})
	})

	r.PATCH("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "user updated! (partial)",
			"id":      id,
		})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "user deleted!",
			"id":      id,
		})
	})

	r.Run(":8888")
}
