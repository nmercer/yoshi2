// TODO: Move this to CMD?
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/namsral/flag"
	"github.com/nmercer/yoshi2/services/server/controller"
	"github.com/nmercer/yoshi2/services/server/handler"
	"github.com/nmercer/yoshi2/services/server/store"
	"github.com/nmercer/yoshi2/services/server/telemetry"
	"google.golang.org/grpc"
)

var (
	tlsCert = flag.String(
		"tls_cert",
		"/var/secrets/tls/server.crt",
		"path to tls cert file")
	tlsKey = flag.String(
		"tls_key",
		"/var/secrets/tls/server.key",
		"path to tls key file")
	tlsEnabled = flag.Bool(
		"tls_enabled",
		false,
		"enable or disable TLS")
	grpcPort = flag.Int(
		"grpc_port",
		50051,
		"port for gRPC")
)

func main() {
	log.Printf("~~ Starting Server on port %d", *grpcPort)

	flag.Parse()

	// setup net listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create server instances
	tempStore := store.NewTempStore()
	tempController := controller.NewTempController(tempStore)
	tempServer := handler.NewTempServer(tempController)

	locationStore := store.NewLocationStore()
	locationController := controller.NewLocationController(locationStore)
	locationServer := handler.NewLocationServer(locationController)

	// TODO: Get TLS working in kubernetes
	// Create the TLS credentials
	// creds, err := credentials.NewServerTLSFromFile(*tlsCert, *tlsKey)
	// if err != nil {
	// 	log.Fatalf("could not load TLS keys: %s", err)
	// }
	// opts := []grpc.ServerOption{grpc.Creds(creds)}

	// // create a gRPC server object
	// grpcServer := grpc.NewServer(opts...)

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	telemetry.RegisterTempsServer(grpcServer, tempServer)
	telemetry.RegisterLocationsServer(grpcServer, locationServer)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
