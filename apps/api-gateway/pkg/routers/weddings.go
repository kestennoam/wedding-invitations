package routers

import (
	"api-gateway/interfaces"
	"api-gateway/pkg/kafka"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"time"
	"wedding-service/proto/gen/gen"
)

type WeddingsRouter struct {
	r Router
}

type weddingHandler struct {
	weddingServiceClient gen.WeddingServiceClient
}

func NewWeddingsRouter(r Router) *WeddingsRouter {
	wr := &WeddingsRouter{r: r}
	return wr
}

func (wr *WeddingsRouter) setupRoutes(router *gin.Engine) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// todo add close connection
	h := &weddingHandler{weddingServiceClient: gen.NewWeddingServiceClient(conn)}

	// TODO extract this outside of this function
	weddingsRouter := router.Group("/weddings")
	{
		weddingsRouter.POST("/", h.createWedding)

	}
}

func (h *weddingHandler) createWedding(c *gin.Context) {
	fmt.Println("lior")
	var wedding interfaces.WeddingDTO
	if err := c.BindJSON(&wedding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wedding.CreatedAt = time.Now()
	wedding.UpdatedAt = time.Now()

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// send with kafka the new wedding
	weddingResponse, err := h.weddingServiceClient.CreateWedding(ctx, &gen.CreateWeddingRequest{
		Wedding: &gen.Wedding{
			Id: wedding.ID,
			Couple: &gen.Couple{
				Person_1: &gen.Person{
					FirstName: wedding.Couple.Person1.FirstName,
					LastName:  wedding.Couple.Person1.LastName,
					Gender:    convertGenderDTOToGender(wedding.Couple.Person1.Gender),
				},
				Person_2: &gen.Person{
					FirstName: wedding.Couple.Person2.FirstName,
					LastName:  wedding.Couple.Person2.LastName,
					Gender:    convertGenderDTOToGender(wedding.Couple.Person2.Gender),
				},
			},
			Date:       timestamppb.New(wedding.Date),
			Location:   wedding.Location,
			GuestCount: int32(wedding.GuestCount),
			CreatedAt:  timestamppb.New(wedding.CreatedAt),
			UpdatedAt:  timestamppb.New(wedding.UpdatedAt),
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//sendWeddingToWeddingsServiceWithREST(wedding)
	fmt.Println("noam")
	//sendWeddingToWeddingsServiceWithKafka(wedding)
	// TODO change to grpc
	c.JSON(201, gin.H{
		"message":    "create wedding",
		"wedding_id": weddingResponse.GetWedding().GetId(),
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

func convertGenderDTOToGender(genderDTO interfaces.GenderDTO) gen.Gender {
	// Implement the conversion logic here
	// This is a placeholder implementation
	switch genderDTO {
	case interfaces.Male:
		return gen.Gender_MALE
	case interfaces.Female:
		return gen.Gender_FEMALE
	default:
		return gen.Gender_OTHER
	}
}
