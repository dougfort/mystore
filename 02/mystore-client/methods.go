package main

import (
	"io"
	"time"

	context "golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/mystore/02/protobuf"
)

type storeClientState struct {
	pb.StoreClient
}

func (s storeClientState) getCatalog() ([]pb.CatalogResponse, error) {
	var psc pb.Store_CatalogStreamClient
	var catalog []pb.CatalogResponse
	var err error

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	if psc, err = s.CatalogStream(ctx, &pb.CatalogStreamRequest{}); err != nil {
		return nil, errors.Wrap(err, "client.CatalogStream")
	}

	for loop := true; loop; {
		var cr *pb.CatalogResponse

		if cr, err = psc.Recv(); err != nil {
			if err == io.EOF {
				loop = false
			} else {
				return nil, errors.Wrap(err, "psc.Recv()")
			}
		} else {
			catalog = append(catalog, *cr)
		}

	}

	return catalog, nil
}

func (s storeClientState) orderItem(userName string, itemID string) error {
	var response *pb.OrderResponse
	var err error

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	request := pb.OrderItemRequest{ItemID: itemID, CustomerName: userName}
	if response, err = s.OrderItem(ctx, &request); err != nil {
		return errors.Wrap(err, "client.OrderItem")
	}

	if !response.Successful {
		return errors.Errorf("unable to order %s", itemID)
	}

	return nil

}
