package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"
)

func getShip(client pb.PlanetExpressClient) (pb.Ship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetShip(ctx, &empty.Empty{})

	if err != nil {
		log.Printf("%v.GetShip(_) = _, %v: ", client, err)
		return pb.Ship{}, err
	}

	return *resp.GetShip(), nil
}

func getCrew(client pb.PlanetExpressClient) (pb.Crew, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetCrew(ctx, &empty.Empty{})

	if err != nil {
		log.Printf("%v.GetCrew(_) = _, %v: ", client, err)
		return pb.Crew{}, err
	}

	return *resp.GetCrew(), nil
}

func listDeliveries(client pb.PlanetExpressClient) ([]*pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.ListDeliveries(ctx, &empty.Empty{})

	if err != nil {
		log.Printf("%v.ListDeliveries(_) = _, %v: ", client, err)
		return []*pb.Delivery{}, err
	}

	return resp.GetDeliveries(), nil
}

func getDelivery(client pb.PlanetExpressClient, getDeliveryRequest *pb.GetDeliveryRequest) (pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetDelivery(ctx, getDeliveryRequest)

	if err != nil {
		log.Printf("%v.GetDelivery(_) = _, %v: ", client, err)
		return pb.Delivery{}, err
	}

	return *resp.GetDelivery(), nil
}

func getWeapons(client pb.PlanetExpressClient) ([]*pb.Weapon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetWeapons(ctx, &empty.Empty{})

	if err != nil {
		log.Printf("%v.GetWeapons(_) = _, %v: ", client, err)
		return []*pb.Weapon{}, err
	}

	return resp.GetWeapons(), nil
}
