package main

import (
	"log"
	"math/rand"

	"google.golang.org/grpc"

	pb "github.com/dougfort/mystore/03/protobuf"
)

func main() {
	const serverAddress = "localhost:1111"
	const userName = "test user"

	var conn *grpc.ClientConn
	var catalog []pb.CatalogResponse
	var err error

	if conn, err = grpc.Dial(serverAddress, grpc.WithInsecure()); err != nil {
		log.Fatalf("grpc.Dial(%s) failed: %v", serverAddress, err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Printf("conn.Close() returned %v", err)
		}
	}()

	client := storeClientState{pb.NewStoreClient(conn)}

	if catalog, err = client.getCatalog(); err != nil {
		log.Fatalf("getCatalog: %v", err)
	}

	log.Printf("found %d catalog entries", len(catalog))

	// choose a random catalog entry to order
	entry := catalog[rand.Int()%len(catalog)]

	log.Printf("ordering (%s) %s at $%d; %d available",
		entry.ItemID,
		entry.Description,
		entry.Price,
		entry.Available,
	)

	if err = client.orderItem(userName, entry.ItemID); err != nil {
		log.Fatalf("orderItem(%s, %s) failed: %s", userName, entry.ItemID, err)
	}

	log.Println("order successful")
}
