package main

import (
    "context"
    "fmt"
    "io"
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

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
    log.Printf("LongGreet function was invoked with a streaming request\n")
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.GreetResponse{
                Result: "EOF message received",
            })
        } else if err != nil {
            log.Fatalf("Error while reading client stream: %v", err)
            return err
        }
        log.Printf("Hello %s %s\n", req.FirstName, req.LastName)
    }
    return nil
}
