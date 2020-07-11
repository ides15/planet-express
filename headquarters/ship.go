package main

import (
	"context"
	"net/http"

	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"
)

type shipResolver struct {
	s pb.Ship
}

func (r *resolver) Ship(ctx context.Context) (*shipResolver, error) {
	ship, err := getShip(client)
	if err != nil {
		return nil, notFoundError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	}

	return &shipResolver{
		s: ship,
	}, nil
}

func (s *shipResolver) Name(ctx context.Context) string {
	return s.s.GetName()
}

func (s *shipResolver) Location(ctx context.Context) string {
	return s.s.GetLocation()
}

type fuelLevelResolver struct {
	fl pb.Ship_FuelLevel
}

func (s *shipResolver) FuelLevel(ctx context.Context) pb.Ship_FuelLevel {
	return s.s.GetFuelLevel()
}

func (s *shipResolver) Crew(ctx context.Context) *pb.Crew {
	return s.s.GetCrew()
}

func (s *shipResolver) Delivery(ctx context.Context) *deliveryResolver {
	return &deliveryResolver{
		d: *s.s.GetDelivery(),
	}
}
