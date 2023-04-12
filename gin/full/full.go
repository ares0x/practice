package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello," + name,
		})
	})
	v1 := router.Group("/api/v1", Auth())
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "list all users",
			})
		})
		v1.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{
				"message": "get user " + id,
			})
		})
	}
	router.Run(":8080")
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 鉴权逻辑
		c.Next()
	}
}
