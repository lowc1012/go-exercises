package main

import (
    "log"
    "net"

    pb "github.com/lowc1012/go-exercises/grpc/greet/proto"
    "google.golang.org/grpc"
)

var addr = "localhost:50051"

// Server is used to implement the pb.GreetServiceServer interface
type Server struct {
    pb.GreetServiceServer
}

func main() {
    // Start the server
    lis, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("Failed to listen on: %v", err)
    }

    log.Printf("Server listening on: %v", addr)
    //
    server := grpc.NewServer()

    // Register the server with the protobuf service server interface
    pb.RegisterGreetServiceServer(server, &Server{})

    // Serve the server
    if err := server.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
