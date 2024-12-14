package main

import (
    "context"
    "io"
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
    log.Println("Starting to do a Unary RPC...")
    res, err := c.Greet(context.Background(), &pb.GreetRequest{
        FirstName: "Ryan",
        LastName:  "Lo",
    })

    if err != nil {
        log.Fatalf("Error while calling Greet RPC: %v", err)
    }
    log.Printf("Greetings: %v", res.Result)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
    log.Println("starting to do a server straming RPC")
    s, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
        FirstName: "Ryan",
        LastName:  "Lo",
    })
    if err != nil {
        log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
    }
    // infinite loop
    for {
        res, err := s.Recv() // receive a response from the stream
        // communication should be closed when the server sends the EOF message
        if err == io.EOF {
            // we've reached the end of the stream
            break
        } else if err != nil {
            log.Fatalf("Error while reading the stream: %v", err)
        }

        log.Printf("Response from GreetManyTimes: %v", res.Result)
    }
}
