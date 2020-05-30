package controller

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nmercer/yoshi2/services/server/store"
)

type TempController interface {
	CreateTemp(temp float32, locationId int32) (*empty.Empty, error)
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
