package mongo

import (
	"context"
	"github.com/fadyat/cloud/internal/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBLayer struct {
	client *mongo.Client
}

func NewDBLayer(connection string) (persistence.DatabaseHandler, error) {
	var clientOptions = options.Client().ApplyURI(connection)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &DBLayer{
		client: client,
	}, nil
}

func (db *DBLayer) AddEvent(event persistence.Event) ([]byte, error) {
	return nil, nil
}

func (db *DBLayer) FindEvent(id []byte) (persistence.Event, error) {
	return persistence.Event{}, nil
}

func (db *DBLayer) FindEventByName(name string) (persistence.Event, error) {
	return persistence.Event{}, nil
}

func (db *DBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	return []persistence.Event{}, nil
}
