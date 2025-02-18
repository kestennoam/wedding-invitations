package main

import (
	"api-gateway/pkg/routers"
)

func main() {

	r := routers.NewRouter()
	r.Run("localhost:8080")

	// Create a new Gin router
	//router := gin.Default()
	//routers.InitRouter(router)
	//// Listen on port 8080
	//router.Run(":8080")
}
