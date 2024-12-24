package main

import (
    "context"
    "io"
    "log"
    "time"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
    stream, err := c.Max(context.Background())

    if err != nil {
        log.Fatalf("Error while calling Max RPC: %v", err)
    }

    ch := make(chan struct{})

    go func() {
        numbers := []int32{
            1, 6, 9, 2, 10, 3,
        }
        log.Printf("Sending numbers: %v\n", numbers)
        for _, n := range numbers {
            err := stream.Send(&pb.MaxRequest{
                Number: n,
            })
            if err != nil {
                log.Fatalf("Error while sending data to server: %v", err)
            }
            time.Sleep(1 * time.Second)
        }
        err := stream.CloseSend()
        if err != nil {
            log.Fatalf("Error while closing the stream: %v", err)
        }
    }()

    var max int32
    go func() {
        for {
            res, err := stream.Recv()
            if err == io.EOF {
                break
            } else if err != nil {
                log.Fatalf("Error while receiving data from server: %v", err)
                break
            }
            max = res.Result
        }
        close(ch)
    }()

    <-ch
    log.Printf("The max number is: %v\n", max)
}
