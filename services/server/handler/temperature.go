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
	return s.controller.CreateTemp(data.Temp, data.LocationId)
}

func (s *tempServer) GetTemps(ctx context.Context, req *telemetry.GetTempsRequest) (*telemetry.GetTempsResponse, error) {
	listTemps, err := s.controller.GetTemp(req.LocationId)
	if err != nil {
		return nil, err
	}

	return &telemetry.GetTempsResponse{
		Temps: listTemps,
	}, nil
}
