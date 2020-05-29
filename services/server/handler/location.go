//TODO: Change name from handler?
package handler

import (
	"log"

	"github.com/nmercer/yoshi2/services/server/telemetry"
	"golang.org/x/net/context"
)

// TempServer represents the gRPC server
type LocationServer struct {
}

func (s *LocationServer) CreateLocation(ctx context.Context, in *telemetry.Location) (*telemetry.Location, error) {
	log.Printf("Name: %s", in.Name)

	// TODO: Make a controller request here
	location := telemetry.Location{Name: in.Name, Id: 1}
	return &location, nil
}
