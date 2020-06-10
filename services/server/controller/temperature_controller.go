package controller

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nmercer/yoshi2/services/server/store"
)

type TempController interface {
	CreateTemp(temp float32, locationId int32) (*empty.Empty, error)
	GetTemp(locationId int32) ([]float32, error)
}

type tempController struct {
	tempStore store.TempStore
}

func NewTempController(tempStore store.TempStore) TempController {
	return &tempController{
		tempStore: tempStore,
	}
}

func (c tempController) CreateTemp(temp float32, locationId int32) (*empty.Empty, error) {
	return c.tempStore.CreateTemp(temp, locationId)
}

func (c tempController) GetTemp(locationId int32) ([]float32, error) {
	return c.tempStore.GetTemp(locationId)
}
