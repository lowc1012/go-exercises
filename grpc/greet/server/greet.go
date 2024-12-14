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

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
    log.Printf("GreetManyTimes function was invoked with: %v\n", req)

    for i := 0; i < 10; i++ {
        err := stream.Send(&pb.GreetResponse{
            Result: fmt.Sprintf("Hello %s %s number %d", req.FirstName, req.LastName, i),
        })
        if err != nil {
            return err
        }
    }
    return nil
}
