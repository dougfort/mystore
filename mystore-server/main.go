package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	pb "github.com/dougfort/mystore/protobuf"
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
	go stopServerOnSignal(grpcServer)

	pb.RegisterStoreServer(grpcServer, NewMyStoreServer())

	log.Printf("Server starts: listening on %s", serverAddress)
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("grpcServer.Serve ended with %s", err)
	}
	log.Printf("Server stops")
}

func stopServerOnSignal(grpcServer *grpc.Server) {
	// set up signals to stop the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s := <-sigChan
	log.Printf("signal: %s", s.String())

	grpcServer.Stop()
}
