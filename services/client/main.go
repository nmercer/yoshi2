package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/namsral/flag"
	"github.com/nmercer/yoshi2/services/client/telemetry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	tlsCert = flag.String(
		"tls_cert",
		"../../server.crt",
		"path to tls cert file")
	tlsEnabled = flag.Bool(
		"tls_enabled",
		false,
		"enable or disable TLS")
	grpcServer = flag.String(
		"grpc_server",
		"192.168.64.3",
		"ip address of grpc server")
	grpcPort = flag.Int(
		"grpc_port",
		31726,
		"port for gRPC")
)

func main() {
	var conn *grpc.ClientConn

	// TODO: Get TLS working in kubernetes
	// 	// Create the client TLS credentials
	// 	creds, err := credentials.NewClientTLSFromFile(*tlsCert, "")
	// 	if err != nil {
	// 		log.Fatalf("could not load tls cert: %s", err)
	// 	}
	// 	conn, err = grpc.Dial(
	// 		fmt.Sprintf("%s:%d", *grpcServer, *grpcPort),
	// 		grpc.WithTransportCredentials(creds))
	// 	if err != nil {
	// 		log.Fatalf("did not connect: %s", err)
	// 	}
	// 	defer conn.Close()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", *grpcServer, *grpcPort),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	locationClient := telemetry.NewLocationsClient(conn)
	locResponse, err := locationClient.CreateLocation(context.Background(), &telemetry.Location{Name: uuid.New().String()})
	if err != nil {
		log.Fatalf("Error when calling CreateLocation: %s", err)
	}

	tempClient := telemetry.NewTempsClient(conn)
	_, err = tempClient.CreateTemp(context.Background(), &telemetry.Temp{Temp: 420.69, LocationId: locResponse.Id})
	if err != nil {
		log.Fatalf("Error when calling CreateTemp: %s", err)
	}
	_, err = tempClient.CreateTemp(context.Background(), &telemetry.Temp{Temp: 69.420, LocationId: locResponse.Id})
	if err != nil {
		log.Fatalf("Error when calling CreateTemp: %s", err)
	}

	tempResponse, err := tempClient.GetTemps(context.Background(), &telemetry.GetTempsRequest{LocationId: locResponse.Id})
	if err != nil {
		log.Fatalf("Error when calling GetTemps: %s", err)
	}

	log.Printf("Location ID: %d", locResponse.Id)
	log.Printf("Location Name: %s", locResponse.Name)
	log.Print(tempResponse.Temps)

	locationsResponse, err := locationClient.GetLocations(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("Error when calling GetLocations: %s", err)
	}
	log.Print(locationsResponse.Locations)
}
