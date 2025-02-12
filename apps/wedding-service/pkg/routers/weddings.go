package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
	"wedding-service/interfaces"
	"wedding-service/pkg/db"
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

	// save to db
	dbInstance, err := db.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to connect to database"})
		return
	}
	defer dbInstance.Close()
	err = dbInstance.InsertWedding(wedding)
	if err != nil {
		log.Printf("failed to insert wedding into table: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to insert wedding into table"})
		return
	}

	log.Printf("inserted wedding into table: %v\n", wedding)
	// send with kafka the new wedding
	c.JSON(201, gin.H{
		"message": "created wedding",
		"wedding": wedding,
	})
}
