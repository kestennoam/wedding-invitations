package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
	pb "wedding-service/proto/gen/gen"
	"wedding-service/pkg/routers"
)

// WeddingServiceServer implements the gRPC server
type WeddingServiceServer struct {
	pb.UnimplementedWeddingServiceServer
	weddings map[string]*pb.Wedding // Simple in-memory storage
}

// NewWeddingServiceServer creates a new instance of WeddingServiceServer
func NewWeddingServiceServer() *WeddingServiceServer {
	return &WeddingServiceServer{
		weddings: make(map[string]*pb.Wedding),
	}
}

// CreateWedding implements the CreateWedding RPC method
func (s *WeddingServiceServer) CreateWedding(ctx context.Context, req *pb.CreateWeddingRequest) (*pb.CreateWeddingResponse, error) {
	log.Printf("Creating wedding: %+v", req)
	if req.Wedding == nil {
		return nil, status.Error(codes.InvalidArgument, "wedding details are required")
	}
	log.Printf("Wedding details: %+v", req.Wedding)


	log.Printf("Validating required fields: %+v", req.Wedding)


	// Store the wedding
	weddingID := generateWeddingID() // You'll need to implement this
	req.Wedding.Id = weddingID
	s.weddings[weddingID] = req.Wedding

	log.Printf("Created wedding with ID: %s", weddingID)

	return &pb.CreateWeddingResponse{
		Wedding: req.Wedding,
	}, nil
}



// Helper function to generate a unique wedding ID
func generateWeddingID() string {
	// For now, we'll use a timestamp-based ID
	return fmt.Sprintf("WED_%d", time.Now().UnixNano())
}

func main() {
	// Create gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	weddingService := NewWeddingServiceServer() // Use the constructor
	pb.RegisterWeddingServiceServer(grpcServer, weddingService)
	
	// Start gRPC server in a goroutine
	go func() {
		log.Printf("Starting gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Create a new Gin router
	router := gin.Default()
	routers.InitRouter(router)
	router.Run("localhost:8081")

	//topic := "create_wedding"
	//wedding, err := kafka.ConsumeWeddingMessage(topic)
	//if err != nil {
	//	log.Fatalf("Failed to consume wedding message: %v", err)
	//}
	//log.Printf("Consumed wedding message: %+v", wedding)
}
