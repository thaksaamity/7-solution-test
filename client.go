package main

import (
	"context"
	"log"
	"time"
	"github.com/thaksananan-01/7-solution-test/beefpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBeefServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetSummary(ctx, &pb.SummaryRequest{})
	if err != nil {
		log.Fatalf("could not fetch summary: %v", err)
	}

	log.Printf("Beef Summary: %v", response.Beef)
}