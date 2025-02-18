package main

import (
	"api-gateway/pkg/routers"
	"context"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWeddingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r := routers.NewRouter()
	r.Run("localhost:8080")

	// Create a new Gin router
	//router := gin.Default()
	//routers.InitRouter(router)
	//// Listen on port 8080
	//router.Run(":8080")
}
