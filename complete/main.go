package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessie-codes/golang-cat-cafe/complete/cat"
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
	r.GET("/cat/:personality", func(c *gin.Context) {
		accepts := c.Request.Header.Get("Accepts")
		result := cats.GetByPersonality(c.Param("personality"))
		if len(result.List) > 0 {
			if accepts == "application/xml" {
				c.XML(http.StatusOK, result)
				return
			}
			c.JSON(http.StatusOK, result)
			return
		}
		if accepts == "application/xml" {
			c.XML(http.StatusNotFound, gin.H{"error": "Not Found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
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
