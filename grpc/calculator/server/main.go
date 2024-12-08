package main

import (
    "log"
    "net"

    pb "github.com/lowc1012/go-exercises/grpc/calculator/proto"
    "google.golang.org/grpc"
)

var addr = "localhost:50052"

type Server struct {
    pb.CalculatorServiceServer
}

func main() {
    //
    lis, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalln("Failed to listen on: ", addr)
    }

    server := grpc.NewServer()
    pb.RegisterCalculatorServiceServer(server, &Server{})

    if err := server.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
