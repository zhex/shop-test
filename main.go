package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	r.POST("/sales", func(c *gin.Context) {
		var req SaleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sale, err := CreateSale(req)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "one or more products not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusCreated, sale)
	})

	r.Run()
}
