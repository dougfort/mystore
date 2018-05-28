package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/dougfort/mystore/03/protobuf"
)

func main() {
	const serverAddress = "localhost:1111"

	var listener net.Listener
	var grpcServer *grpc.Server
	var err error

	if listener, err = net.Listen("tcp", serverAddress); err != nil {
		log.Fatalf("net.Listen failled: %v", err)
	}

	grpcServer = grpc.NewServer()

	pb.RegisterStoreServer(grpcServer, NewMyStoreServer())

	log.Printf("Server starts: listening on %s", serverAddress)
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("grpcServer.Serve ended with %s", err)
	}
}
