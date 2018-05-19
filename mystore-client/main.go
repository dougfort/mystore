package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	const serverAddress = "localhost:1111"

	var conn *grpc.ClientConn
	var err error

	if conn, err = grpc.Dial(serverAddress, grpc.WithInsecure()); err != nil {
		log.Fatalf("grpc.Dial(%s) failed: %v", serverAddress, err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close() returned %v", err)
		}
	}()

}
