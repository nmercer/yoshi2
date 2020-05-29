package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nmercer/yoshi2/services/server/handler"
	"github.com/nmercer/yoshi2/services/server/telemetry"
	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	tempServer := handler.TempServer{}
	locationServer := handler.LocationServer{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	telemetry.RegisterTempsServer(grpcServer, &tempServer)
	telemetry.RegisterLocationsServer(grpcServer, &locationServer)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
