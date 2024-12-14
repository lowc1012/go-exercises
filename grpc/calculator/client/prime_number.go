package main

import (
    "context"
    "io"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func doPrimeNumberDecomposition(client pb.CalculatorServiceClient) {
    stream, err := client.PrimeNumberDecomposition(context.Background(), &pb.PrimeNumberRequest{
        Number: 120,
    })
    if err != nil {
        log.Fatalf("Error while calling PrimeNumberDecomposition RPC: %v", err)
    }

    for {
        res, err := stream.Recv()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalf("Error while reading the stream: %v", err)
        }
        log.Printf("Response from PrimeNumberDecomposition: %v", res.Result)
    }
}
