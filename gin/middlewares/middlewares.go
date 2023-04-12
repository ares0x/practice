package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()
	g.Use(func(c *gin.Context) {
		// 中间件处理逻辑
		c.Next()
	})
	g.Use(auth())
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	g.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, pass, ok := c.Request.BasicAuth(); !ok || user != "admin" || pass != "123456" {
			c.Header("WWW-Authenticate", `Basic realm="My Realm"`)
			c.AbortWithStatus(401)
		} else {
			c.Next()
		}
	}
}
