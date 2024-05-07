package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	pb "main/proto" // Update this with your protobuf package name

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type server struct {
	pb.MyServiceServer
}

func (s *server) ProcessJSON(ctx context.Context, req *pb.Request) (*pb.Array, error) {
	// Unmarshal the JSON data
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(req.JsonData), &data); err != nil {
		return nil, err
	}

	// Convert the data map to Struct
	resultMap := make(map[string]*structpb.Value)
	for key, value := range data {
		val, err := structpb.NewValue(value)
		if err != nil {
			return nil, err
		}
		resultMap[key] = val
	}

	// Prepare the response array
	response := &pb.Array{
		Msg: []*pb.Response{
			{
				ResultMap: resultMap,
			},
		},
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	log.Println("Server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
