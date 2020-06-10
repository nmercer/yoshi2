package controller

import (
	"github.com/nmercer/yoshi2/services/server/store"
	"github.com/nmercer/yoshi2/services/server/telemetry"
)

type LocationController interface {
	CreateLocation(name string) (*telemetry.Location, error)
	GetLocations() ([]*telemetry.Location, error)
}

type locationController struct {
	locationStore store.LocationStore
}

func NewLocationController(locationStore store.LocationStore) LocationController {
	return &locationController{
		locationStore: locationStore,
	}
}

func (c locationController) CreateLocation(name string) (*telemetry.Location, error) {
	return c.locationStore.CreateLocation(name)
}

func (c locationController) GetLocations() ([]*telemetry.Location, error) {
	return c.locationStore.GetLocations()
}
