package main

import (
	"first_proto/api_pb_host"
	"first_proto/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	// Create a new instance of your HostServer
	hostServer := server.NewHostsClient()

	// Register your HostServer with the gRPC server
	api_pb_host.RegisterHostsServer(grpcServer, hostServer)

	// Specify the port on which the server will listen
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the gRPC server
	fmt.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
