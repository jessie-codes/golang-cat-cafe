package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessie-codes/golang-cat-cafe/tutorial/cat"
)

func main() {
	r := gin.Default()
	cats := cat.Get()
	r.GET("/cat", func(c *gin.Context) {
		accepts := c.Request.Header.Get("Accepts")
		if accepts == "application/xml" {
			c.XML(http.StatusOK, cats)
			return
		}
		c.JSON(http.StatusOK, cats)
	})
	r.PATCH("/cat/reserve", func(c *gin.Context) {
		accepts := c.Request.Header.Get("Accepts")
		result := cats.Reserve()
		if result != nil {
			if accepts == "application/xml" {
				c.XML(http.StatusOK, result)
				return
			}
			c.JSON(http.StatusOK, result)
			return
		}
		if accepts == "application/xml" {
			c.XML(http.StatusConflict, gin.H{"error": "Resource Exhausted"})
			return
		}
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Resource Exhausted"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
