package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const port = 10000

var (
	now        = time.Now()
	deliveries = []*pb.Delivery{
		{
			Id:        "99c5e852-7793-42bf-ad2a-0101435c8bfc",
			Timestamp: now.Unix(), // now
			Name:      "Beer",
			Weight:    55,
			Fragile:   false,
		},
		{
			Id:        "2d9243f6-bd59-4b17-a3b3-4e3ee1479856",
			Timestamp: now.AddDate(0, 0, -10).Unix(), // 10 days ago
			Name:      "Glasses",
			Weight:    30,
			Fragile:   true,
		},
		{
			Id:        "ff614c03-99c5-417c-8c23-46e30ae2f822",
			Timestamp: now.AddDate(0, 0, -20).Unix(), // 20 days ago
			Name:      "Chips",
			Weight:    29,
			Fragile:   false,
		},
		{
			Id:        "46bdf5f8-676e-4130-a1f1-dbe0882c17f0",
			Timestamp: now.AddDate(0, 0, -30).Unix(), // 30 days ago
			Name:      "Guacamole",
			Weight:    31,
			Fragile:   false,
		},
	}
	members = []*pb.Member{
		{
			Id:        "355bc0f1-f1ea-49ce-975f-3a0aee96ba32",
			FirstName: "John",
			LastName:  "Ide",
			Age:       24,
		},
		{
			Id:        "679e04f3-4bd2-47ce-8ca9-1ad4952077ce",
			FirstName: "Elon",
			LastName:  "Musk",
			Age:       49,
		},
		{
			Id:        "7e49c347-a44d-41df-b59e-71d74921988b",
			FirstName: "Ken",
			LastName:  "Block",
			Age:       52,
		},
	}
)

type planetExpressShipServer struct {
	pb.UnimplementedPlanetExpressServer
}

func newServer() *planetExpressShipServer {
	s := &planetExpressShipServer{}
	return s
}

func (s *planetExpressShipServer) GetShip(ctx context.Context, empty *empty.Empty) (*pb.GetShipResponse, error) {
	delivery := &pb.Delivery{}

	// Get current delivery
	for _, d := range deliveries {
		if d.Timestamp == now.Unix() {
			delivery = d
		}
	}

	return &pb.GetShipResponse{
		Ship: &pb.Ship{
			Name:      "planet express ship",
			Location:  "omicron persei 8",
			FuelLevel: pb.Ship_FULL,
			Crew: &pb.Crew{
				Size:            int64(len(members)),
				Mood:            pb.Crew_MUTINOUS,
				CaptainMarooned: true,
				Members:         members,
			},
			Delivery: delivery,
		},
	}, nil
}

func (s *planetExpressShipServer) GetCrew(ctx context.Context, empty *empty.Empty) (*pb.GetCrewResponse, error) {
	return &pb.GetCrewResponse{
		Crew: &pb.Crew{
			Size:            int64(len(members)),
			Mood:            pb.Crew_MUTINOUS,
			CaptainMarooned: true,
			Members:         members,
		},
	}, nil
}

func main() {
	log.Println("Planet Express Ship")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPlanetExpressServer(grpcServer, newServer())

	log.Printf("Ship listening on localhost:%d\n", port)
	grpcServer.Serve(lis)
}
