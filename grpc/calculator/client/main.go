package main

import (
    "log"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50052"

func main() {
    // create a connection to the server
    conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to listen on: %v", err)
    }

    // create a client instance
    client := pb.NewCalculatorServiceClient(conn)

    // call the function
    doSum(client)
    doPrimeNumberDecomposition(client)
}
