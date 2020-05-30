package main

import (
	"log"

	"github.com/namsral/flag"
	"github.com/nmercer/yoshi2/services/client/telemetry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tlsCert = flag.String(
		"tls_cert",
		"../../server.crt",
		"path to tls cert file")
)

func main() {
	var conn *grpc.ClientConn

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile(*tlsCert, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	tempClient := telemetry.NewTempsClient(conn)
	_, err = tempClient.CreateTemp(context.Background(), &telemetry.Temp{Temp: 420.69, LocationId: 3})
	if err != nil {
		log.Fatalf("Error when calling CreateTemp: %s", err)
	}

	locationClient := telemetry.NewLocationsClient(conn)
	locResponse, err := locationClient.CreateLocation(context.Background(), &telemetry.Location{Name: "Living Room"})
	if err != nil {
		log.Fatalf("Error when calling CreateLocation: %s", err)
	}

	log.Printf("Location ID: %d", locResponse.Id)
	log.Printf("Location Name: %s", locResponse.Name)

}
