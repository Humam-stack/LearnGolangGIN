package belajargolanggin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestResponseHelper(t *testing.T) {
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "999" {
			ResponseNotFound(c, "User not found")
			return
		}

		ResponseSuccess(c, "User Found", gin.H{
			"id":   id,
			"name": "example",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var input struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required"`
		}

		// validation
		if err := c.ShouldBindJSON(&input); err != nil {
			ResponseBadRequest(c, err.Error())
			return
		}

		ResponseCreated(c, "User Created successfully", gin.H{
			"id":    1,
			"name":  input.Name,
			"email": input.Email,
		})
	})

	r.Run(":8888")
}
