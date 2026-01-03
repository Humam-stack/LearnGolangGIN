package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", getProductWithQuery)

	r.Run(":8888")
}

func getProductWithQuery(c *gin.Context) {
	// // Abil query params (string)

	// page := c.Query("page")
	// limit := c.Query("limit")
	search := c.Query("search")

	// setting default value
	pageDefault := c.DefaultQuery("page", "1")
	limitDefault := c.DefaultQuery("limit", "10")

	//convert string to int

	pageInt, _ := strconv.Atoi(pageDefault)
	limitInt, _ := strconv.Atoi(limitDefault)

	c.JSON(200, gin.H{
		"page":   pageInt,
		"limit":  limitInt,
		"search": search,
	})
}
