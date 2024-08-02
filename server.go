package main

import (
	"gin-demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	v1Router := server.Group("/api/v1")
	routes.ProductRouter(v1Router)
	server.Run(":3000")

}
