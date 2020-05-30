package store

import (
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nmercer/yoshi2/services/server/telemetry"
)

type TempStore interface {
	CreateTemp(temp float32, locationId int32) (*empty.Empty, error)
}

type tempStore struct {
}

func NewTempStore() *tempStore {
	return &tempStore{}
}

func (s tempStore) CreateTemp(temp float32, locationId int32) (*empty.Empty, error) {
	log.Printf("Temp: %f", temp)
	log.Printf("Location: %d", locationId)

	// TODO: Make this a DB call instead
	_ = telemetry.Temp{Temp: temp, LocationId: locationId}
	return new(empty.Empty), nil
}
