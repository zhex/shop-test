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
	r.Run()
}
