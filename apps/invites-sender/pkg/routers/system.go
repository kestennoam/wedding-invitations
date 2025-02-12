package routers

import "github.com/gin-gonic/gin"

func ResponsePong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
