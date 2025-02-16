package routers

import (
	"api-gateway/interfaces"
	"api-gateway/pkg/kafka"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type WeddingsRouter struct {
	r Router
}

func NewWeddingsRouter(r Router) *WeddingsRouter {
	wr := &WeddingsRouter{r: r}
	return wr
}

func (wr *WeddingsRouter) setupRoutes(router *gin.Engine) {
	// TODO extract this outside of this function
	weddingsRouter := router.Group("/weddings")
	{
		weddingsRouter.POST("/", createWedding)

	}
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
	sendWeddingToWeddingsServiceWithREST(wedding)
	//sendWeddingToWeddingsServiceWithKafka(wedding)
	// TODO change to grpc
	c.JSON(201, gin.H{
		"message": "create wedding",
		"wedding": wedding,
	})
}

func sendWeddingToWeddingsServiceWithREST(wedding interfaces.WeddingDTO) {
	jsonBytes, err := json.Marshal(wedding)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/weddings", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	log.Println(resp.Status)
}

func sendWeddingToWeddingsServiceWithKafka(wedding interfaces.WeddingDTO) {
	weddingInBytes, err := json.Marshal(wedding)
	if err != nil {
		log.Printf("error while converting wedding from json to bytes")
		return
	}
	kafka.ProduceMessage("create_wedding", weddingInBytes)
}
