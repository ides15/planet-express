package main

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"
)

// Normal tests
// ---
func TestGetShip_Name(t *testing.T) {
	s := newServer()

	resp, _ := s.GetShip(context.Background(), &empty.Empty{})

	name := resp.GetShip().GetName()
	want := "planet express ship"

	if name != want {
		t.Errorf("wanted %s, got %s", want, name)
	}
}

func TestGetShip_Location(t *testing.T) {
	s := newServer()

	resp, _ := s.GetShip(context.Background(), &empty.Empty{})

	location := resp.GetShip().GetLocation()
	want := "omicron persei 8"

	if location != want {
		t.Errorf("wanted %s, got %s", want, location)
	}
}

func TestGetShip_FuelLevel(t *testing.T) {
	s := newServer()

	resp, _ := s.GetShip(context.Background(), &empty.Empty{})

	fuelLevel := resp.GetShip().GetFuelLevel()
	want := pb.Ship_FULL

	if fuelLevel != want {
		t.Errorf("wanted %s, got %s", want, fuelLevel)
	}
}

// Table-driven tests
// ---
func TestGetDelivery(t *testing.T) {
	s := newServer()

	tests := []struct {
		name  string
		id    string
		isNil bool
	}{
		{
			name:  "ID_FOUND",
			id:    "ff614c03-99c5-417c-8c23-46e30ae2f822",
			isNil: false,
		},
		{
			name:  "ID_NOT_FOUND",
			id:    "12345",
			isNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, _ := s.GetDelivery(context.Background(), &pb.GetDeliveryRequest{Id: "1"})

			delivery := resp.GetDelivery()
			if (delivery != nil) && tt.isNil {
				t.Errorf("wanted no delivery, received %+v\n", delivery)
			}
		})
	}
}
