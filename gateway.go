package main

import (
    "context"
    "log"
    "net/http"
    "google.golang.org/grpc"
    pb "github.com/thaksananan-01/7-solution-test/beefpb"
    gw "github.com/thaksananan-01/7-solution-test/beefpb" // Import the generated Gateway code
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
)

func main() {
    // Create a gRPC client connection
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    mux := runtime.NewServeMux()
    err = pb.RegisterBeefServiceHandler(context.Background(), mux, conn)
    if err != nil {
        log.Fatalf("could not register handler: %v", err)
    }

    // Start the HTTP server
    log.Println("Starting HTTP/REST server on port 8080...")
    http.ListenAndServe(":8080", mux)
}
