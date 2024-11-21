package main

import (
	"github.com/wallisz619/fxck-golang/grpc-guide/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

// main sets up a gRPC server and listens on port 9000.
func main() {
	// Create a TCP listener on port 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	// Start serving gRPC requests over the listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
