package main

import (
    "io"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
    log.Println("Max function was invoked with a streaming request")
    var maxNumber int32

    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil
        } else if err != nil {
            log.Fatalf("Error while reading client stream: %v", err)
        }

        if req.Number > maxNumber {
            maxNumber = req.Number
            if err := stream.Send(&pb.MaxResponse{
                Result: maxNumber,
            }); err != nil {
                log.Fatalf("Error while sending data to client: %v", err)
            }
        }
    }
}
