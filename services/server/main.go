// TODO: Move this to CMD?
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/namsral/flag"
	"github.com/nmercer/yoshi2/services/server/controller"
	"github.com/nmercer/yoshi2/services/server/handler"
	"github.com/nmercer/yoshi2/services/server/store"
	"github.com/nmercer/yoshi2/services/server/telemetry"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
	postgresURL = flag.String(
		"postgres_url",
		"postgres://test:test@10.105.102.104:5432/telemetry?sslmode=disable",
		"url to postgres server")
	httpPort = flag.Int(
		"http_port",
		8080,
		"health check http port")
	grpcHTTPPort = flag.Int(
		"grpc_http_port",
		8081,
		"port for grpc over http")
)

func main() {
	flag.Parse()

	// connect to postgres
	postgresConn, err := pgxpool.Connect(context.Background(), *postgresURL)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer postgresConn.Close()

	// setup net listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create server instances
	tempStore := store.NewTempStore(postgresConn)
	tempController := controller.NewTempController(tempStore)
	tempServer := handler.NewTempServer(tempController)

	locationStore := store.NewLocationStore(postgresConn)
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

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	// attach the Ping service to the server
	telemetry.RegisterTempsServer(grpcServer, tempServer)
	telemetry.RegisterLocationsServer(grpcServer, locationServer)

	// grpc gateway
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = telemetry.RegisterLocationsHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf(":%d", *grpcPort), opts)
	if err != nil {
		log.Fatalf("RegisterYourServiceHandlerFromEndpoint failed: %s", err)
	}
	err = telemetry.RegisterTempsHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf(":%d", *grpcPort), opts)
	if err != nil {
		log.Fatalf("RegisterTempsHandlerFromEndpoint failed: %s", err)
	}

	// grpc prometheus
	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())

	// TODO:
	// Better way than these go func's?
	// These go func should probably be functions that can return an error?
	// We want to kill the server on a bad error so pods can restart.
	// Better way than wait groups?

	// grpc http server
	go func() {
		log.Printf("~~ Starting HTTP GRPC Server on port %d", *grpcHTTPPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *grpcHTTPPort), mux))
	}()

	// health server
	// go func() {
	// 	log.Printf("~~ Starting Health Server on port %d", *httpPort)
	// 	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(http.StatusOK)
	// 	})
	// 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil))
	// }()

	// grpc server
	log.Printf("~~ Starting GRPC Server on port %d", *grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
