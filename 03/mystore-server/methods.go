package main

import (
	"context"
	"log"
	"sync"

	"github.com/pkg/errors"

	pb "github.com/dougfort/mystore/03/protobuf"
)

type serverState struct {
	sync.Mutex
	catalog map[string]catalogEntry
}

// NewMyStoreServer returns an object that implements the pb.StoreServer interface
func NewMyStoreServer() pb.StoreServer {
	var s serverState
	s.catalog = make(map[string]catalogEntry)
	for _, entry := range catalog {
		s.catalog[entry.itemID] = entry
	}
	return &s
}

// CatalogStream returns a stream of catalog entries
func (s *serverState) CatalogStream(
	request *pb.CatalogStreamRequest,
	stream pb.Store_CatalogStreamServer,
) error {
	log.Printf("CatalogStream starts")

	s.Lock()
	defer s.Unlock()

	for key, entry := range s.catalog {
		if err := stream.Send(&pb.CatalogResponse{
			ItemID:      entry.itemID,
			Description: entry.description,
			Price:       int32(entry.price),
			Available:   int32(entry.available),
		}); err != nil {
			return errors.Wrapf(err, "stream.Send %s", key)
		}
	}

	log.Printf("CatalogStream returns normally")
	return nil
}

// OrderItem orders a specified item from the catalog
func (s *serverState) OrderItem(
	ctx context.Context,
	request *pb.OrderItemRequest,
) (*pb.OrderResponse, error) {
	var response pb.OrderResponse

	log.Printf("OrderItem: %v", request.ItemID)

	s.Lock()
	defer s.Unlock()

	entry, ok := s.catalog[request.ItemID]

	if !ok || entry.available == 0 {
		log.Printf("ItemID %s is not available", request.ItemID)
		return &response, nil
	}

	entry.available--
	s.catalog[request.ItemID] = entry

	log.Printf("OrderItem: %v successful", request.ItemID)
	response.Successful = true

	return &response, nil
}
