package Router

import (
	controllers "modules/gin/Controllers"
	middlewares "modules/gin/Middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware)
	admin := router.Group("/admin")
	{
		admin.GET("/", controllers.TestInsert)
	}
	router.Run(":8081")

}
