package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/thaksananan-01/7-solution-test/question3"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
    pb "github.com/thaksananan-01/7-solution-test/beefpb"
)

type BeefServer struct{}

func (s *BeefServer) GetSummary(ctx context.Context, req *pb.SummaryRequest) (*pb.SummaryResponse, error) {
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
	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
			log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBeefServiceServer(grpcServer, &BeefServer{})
	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
}