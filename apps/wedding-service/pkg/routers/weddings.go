package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wedding-service/interfaces"
)

func InitWeddingsRouter(router *gin.Engine) {
	weddingsRouter := router.Group("/weddings")
	{
		weddingsRouter.POST("/", createWedding)
	}
}

func createWedding(c *gin.Context) {
	var wedding interfaces.Wedding
	// convert weddingDTO to wedding
	if err := c.BindJSON(&wedding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wedding.CreatedAt = time.Now()
	wedding.UpdatedAt = time.Now()

	// send with kafka the new wedding
	c.JSON(201, gin.H{
		"message": "created wedding",
		"wedding": wedding,
	})

}
