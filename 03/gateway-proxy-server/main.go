package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	pb "github.com/dougfort/mystore/03/protobuf"
)

func main() {
	const serverAddress = "localhost:1111"
	const proxyAddress = "localhost:8080"
	var listener net.Listener
	var err error

	mux := runtime.NewServeMux()
	ctx := context.Background()

	log.Printf("Proxying to %s", serverAddress)
	err = pb.RegisterStoreHandlerFromEndpoint(
		ctx,
		mux,
		serverAddress,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("RegisterStoreHandlerFromEndpoint failed: %s", err)
	}

	log.Printf("listening to %s", proxyAddress)
	if listener, err = net.Listen("tcp", proxyAddress); err != nil {
		log.Fatalf("net.Listen failed: %s", err)
	}

	http.Serve(listener, mux)
}
