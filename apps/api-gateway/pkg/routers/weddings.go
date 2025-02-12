package routers

import (
	"api-gateway/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitWeddingsRouter(router *gin.Engine) {
	weddingsRouter := router.Group("/weddings")
	{
		weddingsRouter.GET("/:id", getWeddingByID)
		weddingsRouter.POST("/", createWedding)
	}
}

func getWeddingByID(c *gin.Context) {
	// mock data
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":      id,
		"message": "id is " + id,
	})
}

func createWedding(c *gin.Context) {
	var wedding interfaces.WeddingDTO
	if err := c.BindJSON(&wedding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wedding.CreatedAt = time.Now()
	wedding.UpdatedAt = time.Now()

	// send with kafka the new wedding

	c.JSON(201, gin.H{
		"message": "create wedding",
		"wedding": wedding,
	})
}
