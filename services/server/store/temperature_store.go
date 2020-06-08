package store

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TempStore interface {
	CreateTemp(temp float32, locationId int32) (*empty.Empty, error)
}

type tempStore struct {
	postgresConn *pgxpool.Pool
}

func NewTempStore(postgresConn *pgxpool.Pool) *tempStore {
	return &tempStore{
		postgresConn: postgresConn,
	}
}

func (s tempStore) CreateTemp(temp float32, locationId int32) (*empty.Empty, error) {
	createSQL := fmt.Sprintf("INSERT INTO temperatures (temperature, location_id) VALUES ('%f', '%d');", temp, locationId)

	_, err := s.postgresConn.Query(context.Background(), createSQL)
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}
