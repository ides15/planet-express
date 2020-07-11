package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"

	"google.golang.org/grpc"
)

const planetExpressFilename = "./planet_express.json"

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	client     = pb.NewPlanetExpressClient(nil)
)

type planetExpressData struct {
	Ship          pb.Ship        `json:"ship"`
	AllDeliveries []*pb.Delivery `json:"all_deliveries"`
	Crew          pb.Crew        `json:"crew"`
	Weapons       []*pb.Weapon   `json:"weapons"`
}

func generatePlanetExpressJSON() {
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

	weapons, _ := getWeapons(client)
	log.Println("WEAPONS:")
	for _, weapon := range weapons {
		log.Printf("%+v\n", weapon)
	}
	log.Println()

	// Take planet express information and write to file
	data := &planetExpressData{
		Ship:          ship,
		Crew:          crew,
		AllDeliveries: deliveries,
		Weapons:       weapons,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	if err := ioutil.WriteFile(planetExpressFilename, file, 0644); err != nil {
		log.Fatal("could not write planet express data to file")
	}

	log.Printf("%s created\n", planetExpressFilename)
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

	client = pb.NewPlanetExpressClient(conn)
	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)

	generatePlanetExpressJSON()

	// GraphQL
	StartServer()
}
