package mongo

import (
	"context"
	"fmt"
	"github.com/fadyat/cloud/internal/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBLayer struct {
	client *mongo.Client
}

func NewDBLayer(connection string) (*DBLayer, error) {
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
	s, err := db.client.StartSession()
	fmt.Println("Started session")
	if err != nil {
		return nil, err
	}
	defer s.EndSession(context.Background())

	insert := func(sc mongo.SessionContext) (interface{}, error) {
		result, e := db.client.Database("cloud").
			Collection("events").
			InsertOne(sc, event)

		fmt.Printf("Inserted event: %+v, %+v\n", result, err)

		if e != nil {
			return nil, e
		}

		return result.InsertedID, nil
	}

	id, err := s.WithTransaction(context.Background(), insert)
	fmt.Printf("Transaction event: %+v, %+v\n", id, err)
	if err != nil {
		return nil, err
	}

	return id.([]byte), nil
}

func (db *DBLayer) FindEvent(id []byte) (persistence.Event, error) {
	panic("implement me")
}

func (db *DBLayer) FindEventByName(name string) (persistence.Event, error) {
	panic("implement me")
}

func (db *DBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	panic("implement me")
}
