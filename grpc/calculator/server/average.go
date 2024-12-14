package main

import (
    "io"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
    log.Println("Average function was invoked with a streaming request")
    sum := int32(0)
    count := 0
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.AvgResponse{
                Result: float32(sum / int32(count)),
            })
        } else if err != nil {
            log.Fatalf("Error while reading client stream: %v", err)
            return err
        }
        log.Printf("Received number: %v\n", req)
        sum += req.Number
        count++
    }
    return nil
}
