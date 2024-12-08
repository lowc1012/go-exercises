package main

import (
    "context"
    "fmt"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
    log.Printf("Greet function was invoked with: %v\n", req)
    return &pb.GreetResponse{
        Result: fmt.Sprintf("Hello %s %s", req.FirstName, req.LastName),
    }, nil
}
