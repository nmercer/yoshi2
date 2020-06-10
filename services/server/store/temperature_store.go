package store

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TempStore interface {
	CreateTemp(temp float32, locationId int32) (*empty.Empty, error)
	GetTemp(locationId int32) ([]float32, error)
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

func (s tempStore) GetTemp(locationId int32) ([]float32, error) {
	// TODO: How do I make sure this is sorted by date "created"
	getSQL := fmt.Sprintf("SELECT temperature FROM temperatures WHERE location_id=%d", locationId)

	rows, err := s.postgresConn.Query(context.Background(), getSQL)
	if err != nil {
		return nil, err
	}

	var temps []float32

	// TODO: There has to be a better way to do this
	for rows.Next() {
		var temp float32

		err := rows.Scan(&temp)
		if err != nil {
			return nil, err
		}

		temps = append(temps, temp)
	}

	return temps, nil
}
