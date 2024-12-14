package main

import (
    "context"
    "log"
    "time"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func doAverage(client pb.CalculatorServiceClient) {
    log.Println("Starting to do a client streaming RPC...")
    s, err := client.Average(context.Background())
    if err != nil {
        log.Fatalf("Error while calling Average RPC: %v", err)
    }
    reqs := []*pb.AvgRequest{
        {
            Number: 1,
        },
        {
            Number: 2,
        },
        {
            Number: 3,
        },
        {
            Number: 4,
        },
        {
            Number: 5,
        },
    }

    for _, req := range reqs {
        log.Printf("Sending req: %v\n", req)
        if err := s.Send(req); err != nil {
            log.Fatalf("Error while sending req: %v", err)
        }
        time.Sleep(1000 * time.Millisecond)
    }

    res, err := s.CloseAndRecv()
    if err != nil {
        log.Fatalf("Error while receiving response from Average: %v", err)
    }
    log.Printf("Average Response: %v", res.Result)
}
