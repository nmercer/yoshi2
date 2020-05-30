//TODO: Change name from handler?
package handler

import (
	"github.com/nmercer/yoshi2/services/server/controller"
	"github.com/nmercer/yoshi2/services/server/telemetry"
	"golang.org/x/net/context"
)

// TempServer represents the gRPC server
type locationServer struct {
	controller controller.LocationController
}

func NewLocationServer(controller controller.LocationController) telemetry.LocationsServer {
	return &locationServer{
		controller: controller,
	}
}

func (s *locationServer) CreateLocation(ctx context.Context, data *telemetry.Location) (*telemetry.Location, error) {
	// TODO: Pass GRPC or raw data here like name?
	return s.controller.CreateLocation(data.Name)
}
