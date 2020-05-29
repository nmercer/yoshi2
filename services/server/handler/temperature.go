package handler

import (
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nmercer/yoshi2/services/server/telemetry"
	"golang.org/x/net/context"
)

// TempServer represents the gRPC server
type TempServer struct {
}

func (s *TempServer) CreateTemp(ctx context.Context, in *telemetry.Temp) (*empty.Empty, error) {
	// TODO: Make a controller request here
	log.Printf("Temp: %f", in.Temp)
	log.Printf("Location: %d", in.LocationId)

	return new(empty.Empty), nil
}
