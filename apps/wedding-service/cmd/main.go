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
	router.Run("localhost:8081")

	//topic := "create_wedding"
	//wedding, err := kafka.ConsumeWeddingMessage(topic)
	//if err != nil {
	//	log.Fatalf("Failed to consume wedding message: %v", err)
	//}
	//log.Printf("Consumed wedding message: %+v", wedding)
}
