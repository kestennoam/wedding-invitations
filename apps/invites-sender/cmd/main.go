package main

import (
	"github.com/gin-gonic/gin"
	"invites-sender/pkg/routers"
)

func main() {

	r := gin.Default()
	r.GET("/ping", routers.ResponsePong)
	r.Run(":8080")
}
