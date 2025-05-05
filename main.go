package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(errorHandlingMiddleware())

	r.GET("/products", func(c *gin.Context) {
		products, err := GetAllProducts()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, products)
	})

	r.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := CreateProduct(&product); err != nil {
			panic(err)
		}
		c.JSON(http.StatusCreated, product)
	})

	r.Run()
}
