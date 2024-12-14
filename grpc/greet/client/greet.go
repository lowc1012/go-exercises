package main

import (
    "context"
    "io"
    "log"
    "time"

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

func doLongGreet(c pb.GreetServiceClient) {
    log.Println("Starting to do a LongGreet RPC...")
    reqs := []*pb.GreetRequest{
        {
            FirstName: "Ryan",
            LastName:  "Lo",
        },
        {
            FirstName: "Steph",
            LastName:  "Curry",
        },
        {
            FirstName: "Klay",
            LastName:  "Thompson",
        },
    }
    stream, err := c.LongGreet(context.Background())
    if err != nil {
        log.Fatalf("Error while calling LongGreet RPC: %v", err)
    }

    for _, req := range reqs {
        log.Printf("Sending req %v\n", req)
        if err := stream.Send(req); err != nil {
            log.Fatalf("Error while sending to server: %v", err)
        }
        time.Sleep(1000 * time.Millisecond)
    }

    res, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("Error while receiving response from LongGreet: %v", err)
    }
    log.Printf("LongGreet Response: %v", res.Result)
}
