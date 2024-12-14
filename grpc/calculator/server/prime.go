package main

import (
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func (s *Server) PrimeNumberDecomposition(req *pb.PrimeNumberRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
    log.Printf("PrimeNumberDecomposition function was invoked with: %v\n", req)
    num := req.Number
    div := int32(2)
    for num > 1 {
        if num%div == 0 {
            log.Printf("Divisor found: %v\n", div)
            if err := stream.Send(&pb.PrimeNumberResponse{
                Result: div,
            }); err != nil {
                return err
            }
            num = num / div
        } else {
            div++
            log.Printf("Divisor has increased to: %v\n", div)
        }
    }
    return nil
}
