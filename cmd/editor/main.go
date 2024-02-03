package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Nagzham/grpc_host_server/internal/server"
	"github.com/Nagzham/grpc_host_server/pkg/api/editor"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	// Create a new instance of your HostServer
	hostServer := server.NewHostsServer()

	// Register your HostServer with the gRPC server
	editor.RegisterQuizServiceServer(grpcServer, hostServer)

	// Specify the port on which the server will listen
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the gRPC server
	fmt.Println("Starting gRPC server on port 50051...")
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
