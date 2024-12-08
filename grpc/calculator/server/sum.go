package main

import (
    "context"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
    log.Println("Sum function was invoked with: ", req)
    sum := req.FirstNumber + req.SecondNumber
    return &pb.SumResponse{
        Result: sum,
    }, nil
}
