package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const planetExpressFilename = "../planet_express.json"

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

type planetExpressData struct {
	Ship          pb.Ship        `json:"ship"`
	AllDeliveries []*pb.Delivery `json:"all_deliveries"`
	Crew          pb.Crew        `json:"crew"`
}

func getShip(client pb.PlanetExpressClient) (pb.Ship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetShip(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return pb.Ship{}, err
	}

	return *resp.Ship, nil
}

func getCrew(client pb.PlanetExpressClient) (pb.Crew, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetCrew(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.GetCrew(_) = _, %v: ", client, err)
		return pb.Crew{}, err
	}

	return *resp.Crew, nil
}

func listDeliveries(client pb.PlanetExpressClient) ([]*pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.ListDeliveries(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.ListDeliveries(_) = _, %v: ", client, err)
		return []*pb.Delivery{}, err
	}

	return resp.Deliveries, nil
}

func getDelivery(client pb.PlanetExpressClient, getDeliveryRequest *pb.GetDeliveryRequest) (pb.Delivery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetDelivery(ctx, getDeliveryRequest)

	if err != nil {
		log.Fatalf("%v.GetDelivery(_) = _, %v: ", client, err)
		return pb.Delivery{}, err
	}

	return *resp.Delivery, nil
}

func main() {
	log.Println("Planet Express Headquarters")

	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)
	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)

	ship, _ := getShip(client)
	log.Println("SHIP:")
	log.Printf("%+v\n", ship)
	log.Println()

	crew, _ := getCrew(client)
	log.Println("CREW:")
	log.Println(crew)
	log.Println()

	deliveries, _ := listDeliveries(client)
	log.Println("DELIVERIES:")
	for _, delivery := range deliveries {
		log.Printf("%+v\n", delivery)
	}
	log.Println()

	delivery, _ := getDelivery(client, &pb.GetDeliveryRequest{Id: "ff614c03-99c5-417c-8c23-46e30ae2f822"})
	log.Println("SINGLE DELIVERY:")
	log.Printf("%+v\n", delivery)
	log.Println()

	// Take planet express information and write to file
	data := &planetExpressData{
		Ship:          ship,
		Crew:          crew,
		AllDeliveries: deliveries,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	if err = ioutil.WriteFile(planetExpressFilename, file, 0644); err != nil {
		log.Fatal("could not write planet express data to file")
	}

	log.Printf("%s created\n", planetExpressFilename)
}
