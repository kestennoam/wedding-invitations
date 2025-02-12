package main

import (
	"api-gateway/pkg/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()
	routers.InitRouter(router)
	// Listen on port 8080
	router.Run(":8080")
}
