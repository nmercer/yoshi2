package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nmercer/yoshi2/services/server/telemetry"
)

type LocationStore interface {
	CreateLocation(name string) (*telemetry.Location, error)
	GetLocations() ([]*telemetry.Location, error)
}

type locationStore struct {
	postgresConn *pgxpool.Pool
}

func NewLocationStore(postgresConn *pgxpool.Pool) *locationStore {
	return &locationStore{
		postgresConn: postgresConn,
	}
}

func (s locationStore) CreateLocation(name string) (*telemetry.Location, error) {
	// TODO: Some generic protection against SQL injection here?
	createSQL := fmt.Sprintf("INSERT INTO locations (name) VALUES ('%s') RETURNING location_id;", name)

	var locationID int32
	err := s.postgresConn.QueryRow(context.Background(), createSQL).Scan(&locationID)
	if err != nil {
		return nil, err
	}

	// TODO: Catch and return duplicate key errors here?

	location := telemetry.Location{Name: name, Id: locationID}
	return &location, nil
}

func (s locationStore) GetLocations() ([]*telemetry.Location, error) {
	getSQL := fmt.Sprint("SELECT * FROM locations")

	rows, err := s.postgresConn.Query(context.Background(), getSQL)
	if err != nil {
		return nil, err
	}

	var locations []*telemetry.Location

	for rows.Next() {
		var locationID int32
		var name string

		err := rows.Scan(&locationID, &name)
		if err != nil {
			return nil, err
		}

		location := telemetry.Location{Name: name, Id: locationID}
		locations = append(locations, &location)
	}

	return locations, nil
}
