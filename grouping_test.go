package belajargolanggin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGroupingRoute(t *testing.T) {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", getAllProducts)
			products.POST("", createProduct)
			products.GET("/:id", getProductByID)
			products.PUT("/:id", updateProduct)
			products.PATCH("/:id", patchProduct)
			products.DELETE("/:id", deleteProduct)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("", getAllCategories)
			categories.POST("", createCategory)
			categories.GET("/:id", getCategoryByID)
			categories.DELETE("/:id", deleteCategory)
		}
	}

	r.Run(":8888")
}

func getAllProducts(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "get all product",
		"data": []gin.H{
			{
				"id":    1,
				"name":  "Laptop ASUS",
				"price": 15000000,
				"stock": 10,
			},
			{
				"id":    2,
				"name":  "Mouse Gaming",
				"price": 350000,
				"stock": 25,
			},
		},
	})
}

func createProduct(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "product created",
		"data": gin.H{
			"id":    3,
			"name":  "Produk Baru",
			"price": 100000,
			"stock": 10,
		},
	})
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Get Product by ID",
		"id":      id,
		"data": gin.H{
			"id":    id,
			"name":  "Laptop ASUS",
			"price": 1500000,
			"stock": 10,
		},
	})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "update product sukses",
		"id":      id,
	})
}

func patchProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "patch product sukses",
		"id":      id,
	})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "deleted product sukses",
		"id":      id,
	})
}

func getAllCategories(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get all categories",
		"data": []gin.H{
			{
				"id":   1,
				"name": "Elektronik",
			},
			{
				"id":   2,
				"name": "Fashion",
			},
		},
	})
}

func createCategory(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "category created",
		"data": gin.H{
			"id":   3,
			"name": "Books",
		},
	})
}

func getCategoryByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "get category by ID",
		"id":      id,
		"data": gin.H{
			"id":   id,
			"name": "Elektronik",
		},
	})
}

func deleteCategory(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "sukses delete category",
		"id":      id,
	})
}
