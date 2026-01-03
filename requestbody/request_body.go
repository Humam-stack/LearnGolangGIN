package main

import "github.com/gin-gonic/gin"

type CreateProductInput struct {
	Name  string  `json:"name" binding:"required,min=3"`
	Price float64 `json:"price" bindng:"required,min=0"`
	Stock int     `json:"stock" binding:"required,min=0"`
}

type CreateUser struct {
	Name     string `json:"name" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required"`
	Password string `json:"password" binding:"required,min=5"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.POST("/products", createProductHandler)
	r.POST("/users", createUser)

	r.Run(":8888")
}

func createProductHandler(c *gin.Context) {
	var inputUser CreateProductInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"error":   "validation error",
			"details": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "sukses membuat produk",
		"data":    inputUser,
	})
}

func createUser(c *gin.Context) {
	var input CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"details": err.Error(),
		})
	}

	user := UserResponse{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}

	c.JSON(201, gin.H{
		"success": true,
		"data":    user,
	})
}
