package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

var Grpcaddr := ":9093"

func main() {
	//start gRPC server
	lis, err := net.Listen("tcp, GrpcAddr")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	NewGrpcHandler(grpcServer)

	log.Printf("Starting the grpc server trip-service on port %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}