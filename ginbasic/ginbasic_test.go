package belajargolanggin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinBasic(t *testing.T) {
	// membuat router gin
	r := gin.Default()

	r.GET("/ping", handlePing)
	r.GET("/user", handleNested)
	r.GET("/users", handleSlice)
	r.GET("/hello-me", handleHelloMe)
	r.GET("/morning", handleMoring)
	r.GET("/afternoon", handleAfternoon)
	r.GET("/night", handleNight)
	r.Run(":8888")
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func handleNested(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": gin.H{
			"name":  "humam",
			"email": "admin@gmail.com",
		},
	})
}

func handleSlice(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": []string{"humam", "user3", "user3"},
	})
}

func handleHelloMe(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Humam",
	})
}

func handleMoring(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "good morning",
	})
}

func handleAfternoon(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "good afternoon",
	})
}

func handleNight(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "good night",
	})
}
