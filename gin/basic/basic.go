package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Handle("GET", "/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	g.GET("/hello/basic", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	g.Run(":8080")
}
