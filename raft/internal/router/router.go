package router

import "github.com/gin-gonic/gin"

func LoadEngine() *gin.Engine {
	g := gin.Default()
	loadRouters(g)
	return g
}

// loadMiddlewares 加载中间件
func loadMiddlewares(g *gin.Engine) {

}

func loadRouters(g *gin.Engine) {
	storageHandler := NewStorageHandler()
	storageGroup := g.Group("/storage/v1")
	{
		storageGroup.POST("/set", storageHandler.Set)
		storageGroup.GET("/get", storageHandler.Get)
	}
}
