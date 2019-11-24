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
	r.Run() // listen and serve on 0.0.0.0:8080
}
