package main

import (
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    pb "github.com/lowc1012/go-exercises/grpc/greet/proto"
)

var addr = "localhost:50051"

func main() {
    conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // insecure.NewCredentials() is used for testing purposes only
    if err != nil {
        log.Fatalf("Failed to connect server: %v", err)
    }

    defer conn.Close()

    client := pb.NewGreetServiceClient(conn)

    doGreet(client)
    doGreetManyTimes(client)
}
