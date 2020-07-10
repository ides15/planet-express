package main

import (
	"context"
	"net/http"

	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"
)

type deliveryResolver struct {
	d pb.Delivery
}

func (r *resolver) GetDelivery(ctx context.Context, args struct{ ID string }) (*deliveryResolver, error) {
	delivery, err := getDelivery(client, &pb.GetDeliveryRequest{Id: args.ID})
	if err != nil {
		return nil, deliveryNotFoundError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	}

	return &deliveryResolver{
		d: delivery,
	}, nil
}

func (d *deliveryResolver) ID(ctx context.Context) string {
	return d.d.GetId()
}

func (d *deliveryResolver) Timestamp(ctx context.Context) int32 {
	return int32(d.d.GetTimestamp())
}

func (d *deliveryResolver) Name(ctx context.Context) string {
	return d.d.GetName()
}

func (d *deliveryResolver) Weight(ctx context.Context) int32 {
	return int32(d.d.GetWeight())
}

func (d *deliveryResolver) Fragile(ctx context.Context) bool {
	return d.d.GetFragile()
}
