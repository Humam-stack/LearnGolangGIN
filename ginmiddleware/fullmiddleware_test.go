package belajargolanggin

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestFullMiddleware(t *testing.T) {
	r := gin.Default()

	// global middleware

	r.Use(customLoggerMiddleware())
	r.Use(gin.Recovery())

	//public routes
	r.POST("/login", loginHandlers)

	api := r.Group("/api")
	api.Use(authMiddleware())
	{
		api.GET("/profile", ProfileHandlers)
		api.GET("/users", UserHandlers)
		api.POST("/products", CreateProductHandler)
	}

	r.Run(":8888")

}

// Middleware
func customLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()

		fmt.Printf("Cuustom Log - %s %s | Stats : %d | Duration %v\n", method, path, status, duration)
	}
}

// Handlers
func loginHandlers(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		ResponseBadRequest(c, err.Error())
		return
	}

	// simulasi untk login
	if input.Email == "admin@gmail.com" && input.Password == "password" {
		ResponseSuccess(c, "Sukses Login", gin.H{
			"token": "secret",
		})
		return
	}
	ResponseAuthorized(c, "invalid Credentials")
}

func ProfileHandlers(c *gin.Context) {
	ResponseSuccess(c, "Profile details", gin.H{
		"id":    1,
		"Name":  "admin",
		"Email": "admin@gmail.com",
	})
}

func UserHandlers(c *gin.Context) {
	ResponseSuccess(c, "Users", []gin.H{
		{
			"id":   1,
			"name": "admin",
		},
		{
			"id":   2,
			"name": "admin2",
		},
		{
			"id":   3,
			"name": "admin3",
		},
	})
}

func CreateProductHandler(c *gin.Context) {
	var inputProduct struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required,min=0"`
		Stock int     `json:"stock" binging:"required,min=0"`
	}

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		ResponseBadRequest(c, err.Error())
		return
	}

	ResponseCreated(c, "produk berhasil dibuat", gin.H{
		"id":    1,
		"name":  inputProduct.Name,
		"price": inputProduct.Price,
		"stock": inputProduct.Stock,
	})
}
