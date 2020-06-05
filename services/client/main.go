package main

import (
	"fmt"
	"log"

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
