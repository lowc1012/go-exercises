package main

import (
    "context"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
    // call the function
    res, err := c.Sum(context.Background(), &pb.SumRequest{
        FirstNumber:  3,
        SecondNumber: 10,
    })
    if err != nil {
        log.Fatalf("Error while calling Sum RPC: %v", err)
    }
    log.Printf("Sum: %v", res.Result)
}
