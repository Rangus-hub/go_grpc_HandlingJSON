package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "main/proto" // Update this with your protobuf package name
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewMyServiceClient(conn)

	// Prepare your JSON data
	jsonData := `{"user": 1, "name": "Raju"}`

	//requesting(client, jsonData)

	// Send the request
	response, err := client.ProcessJSON(context.Background(), &pb.Request{JsonData: jsonData})
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	// Print the response
	log.Printf("Response TYPE: %T", response)
	log.Printf("Response: %v", response)
}
