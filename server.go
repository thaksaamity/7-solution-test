package main

import (
	"https://github.com/thaksananan-01/7-solution-test/question3"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"https://github.com/thaksananan-01/7-solution-test/beefpb" // Import the generated protobuf package
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedBeefServiceServer
}

func (s *server) GetSummary(ctx context.Context, req *pb.SummaryRequest) (*pb.SummaryResponse, error) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	text := string(body)
	processedData := question3.PieFireDire(text)

	// Convert the result to the gRPC response format
	result := &pb.SummaryResponse{
		Beef: make(map[string]int32),
	}

	for key, value := range processedData {
		result.Beef[key] = int32(value.(int))
	}

	return result, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBeefServiceServer(grpcServer, &server{})

	// Enable reflection for debugging with gRPC tools
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
