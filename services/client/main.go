package main

import (
	"log"

	"github.com/nmercer/yoshi2/services/client/telemetry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	tempClient := telemetry.NewTempsClient(conn)
	_, err = tempClient.CreateTemp(context.Background(), &telemetry.Temp{Temp: 420.69, LocationId: 3})
	if err != nil {
		log.Fatalf("Error when calling GetTemp: %s", err)
	}

	locationClient := telemetry.NewLocationsClient(conn)
	locResponse, err := locationClient.CreateLocation(context.Background(), &telemetry.Location{Name: "Living Room"})
	if err != nil {
		log.Fatalf("Error when calling CreateLocation: %s", err)
	}

	log.Printf("Location ID: %d", locResponse.Id)
	log.Printf("Location Name: %s", locResponse.Name)
	// response, err := c.SayHello(context.Background(), &telemetry.PingMessage{Greeting: "foo"})

}
