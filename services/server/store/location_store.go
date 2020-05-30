package store

import (
	"github.com/nmercer/yoshi2/services/server/telemetry"
)

type LocationStore interface {
	CreateLocation(name string) (*telemetry.Location, error)
}

type locationStore struct {
}

func NewLocationStore() *locationStore {
	return &locationStore{}
}

func (s locationStore) CreateLocation(name string) (*telemetry.Location, error) {
	// TODO: Make this a DB call instead
	location := telemetry.Location{Name: name, Id: 2}
	return &location, nil
}
