package handler

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nmercer/yoshi2/services/server/controller"
	"github.com/nmercer/yoshi2/services/server/telemetry"
	"golang.org/x/net/context"
)

// TempServer represents the gRPC server
type tempServer struct {
	controller controller.TempController
}

func NewTempServer(controller controller.TempController) telemetry.TempsServer {
	return &tempServer{
		controller: controller,
	}
}

func (s *tempServer) CreateTemp(ctx context.Context, data *telemetry.Temp) (*empty.Empty, error) {
	// TODO: Pass GRPC or raw data here like name?
	return s.controller.CreateTemp(data.Temp, data.LocationId)
}
