package routes

import (
	controllers "gin-demo/Controllers"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.RouterGroup) *gin.RouterGroup {
	productRouter := router.Group("/product")
	{
		productRouter.GET("/:id", controllers.GetProduct)
		productRouter.GET("/all", controllers.GetAllProducts)
		productRouter.POST("/", controllers.CreateProduct)
		// productRouter.DELETE("/")
	}

	return productRouter
}
