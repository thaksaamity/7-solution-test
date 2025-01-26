package main

import (
    "context"
    "log"
    "net"
    pb "github.com/thaksananan-01/7-solution-test/question3" // Update this import to point to your generated package
    "google.golang.org/grpc"
)

type BeefServer struct {
    pb.UnimplementedBeefServiceServer // Embed the unimplemented service to fulfill the interface
}

// Implement the GetSummary method for the BeefServiceServer interface
func (s *BeefServer) GetSummary(ctx context.Context, req *pb.SummaryRequest) (*pb.SummaryResponse, error) {
    // Your business logic here to return a summary
    return &pb.SummaryResponse{
        Beef: map[string]int32{"bacon": 42, "ipsum": 15},
    }, nil
}

func main() {
    // Start gRPC server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()

    // Register the BeefServer as a BeefServiceServer
    pb.RegisterBeefServiceServer(grpcServer, &BeefServer{})

    log.Println("gRPC server running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
