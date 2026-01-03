package belajargolanggin

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func customLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//dilakukan sebelum request
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next() // <== lanjut ke handler

		// setelah request
		duration := time.Since(start)
		status := c.Writer.Status()

		fmt.Printf("CUSTOM LOG - %s %s | Status : %d | Duration : %v\n", method, path, status, duration)
	}
}

func TestCustomLogger(t *testing.T) {
	r := gin.New()

	r.Use(customLogger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello bro!",
		})
	})

	r.Run(":8888")
}
