package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"

    api "github.com/thaksananan-01/7-solution-test/api"
    "github.com/thaksananan-01/7-solution-test/question1"
    "github.com/thaksananan-01/7-solution-test/question2"
    "github.com/thaksananan-01/7-solution-test/question3"
    "github.com/thaksananan-01/7-solution-test/utils"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type APIServer struct {
    api.UnimplementedAPIServiceServer
}

func (s *APIServer) GetSummary(ctx context.Context, req *api.SummaryRequest) (*api.SummaryResponse, error) {
    response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
    if err != nil {
        return nil, fmt.Errorf("failed to fetch data: %v", err)
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }

    text := string(body)
    processedData := question3.PieFireDire(text)

    result := &api.SummaryResponse{Beef: make(map[string]int32)}
    for key, value := range processedData {
        result.Beef[key] = int32(value.(int))
    }

    return result, nil
}

func (s *APIServer) GetTriangleWithLoop(ctx context.Context, req *api.TriangleRequest) (*api.TriangleResponse, error) {
    data, err := utils.ReadJSON("files/numberList.json")
    if err != nil {
        return nil, fmt.Errorf("failed to read JSON: %v", err)
    }
    sum := question1.TriangleSumValueWithLoop(data)
    return &api.TriangleResponse{Sum: int32(sum)}, nil
}

func (s *APIServer) GetTriangleWithNonLoop(ctx context.Context, req *api.TriangleRequest) (*api.TriangleResponse, error) {
    data, err := utils.ReadJSON("files/numberList.json")
    if err != nil {
        return nil, fmt.Errorf("failed to read JSON: %v", err)
    }
    sum := question1.TriangleSumValueWithoutLoop(data)
    return &api.TriangleResponse{Sum: int32(sum)}, nil
}

func (s *APIServer) CatchMeIfYouCan(ctx context.Context, req *api.CatchMeRequest) (*api.CatchMeResponse, error) {
    // Extract the single input string
    input := req.Input

    // Decode the string
    decoded := question2.DecodeString(input)

    // Convert the decoded result to a string (if it returns []int or similar)
    output := fmt.Sprint(decoded)

    // Return the result
    return &api.CatchMeResponse{Output: output}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    api.RegisterAPIServiceServer(grpcServer, &APIServer{})
    reflection.Register(grpcServer)

    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve gRPC: %v", err)
        }
    }()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}

    err = api.RegisterAPIServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
    if err != nil {
        log.Fatalf("failed to register gateway: %v", err)
    }

    log.Println("Starting HTTP server on port 8080...")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatalf("failed to serve HTTP: %v", err)
    }
}
