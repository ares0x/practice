package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()
	apiGroup := g.Group("/api")
	{
		apiGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})
	}
	g.Run(":8080")
}
