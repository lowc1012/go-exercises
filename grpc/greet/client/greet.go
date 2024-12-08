package main

import (
    "context"
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
