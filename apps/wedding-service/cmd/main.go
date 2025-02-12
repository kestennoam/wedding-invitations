package main

import (
	"github.com/gin-gonic/gin"
	"wedding-service/pkg/routers"
)

func main() {
	// Create a new Gin router
	router := gin.Default()
	routers.InitRouter(router)
	// Listen on port 8080
	router.Run(":8081")
}
