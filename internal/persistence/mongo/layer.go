package mongo

import (
	"context"
	"github.com/fadyat/cloud/internal/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (db *DBLayer) CreateEvent(event persistence.Event) ([]byte, error) {
	s, err := db.client.StartSession()
	defer s.EndSession(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := db.client.Database("db").
		Collection("events").
		InsertOne(context.Background(), event)

	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	return id[:], nil
}

func (db *DBLayer) FindEvent(id []byte) (persistence.Event, error) {
	panic("implement me")
}

func (db *DBLayer) FindEventByName(name string) (persistence.Event, error) {
	panic("implement me")
}

func (db *DBLayer) FindAll() ([]persistence.Event, error) {
	s, err := db.client.StartSession()
	defer s.EndSession(context.Background())
	if err != nil {
		return nil, err
	}

	cursor, err := db.client.Database("db").
		Collection("events").
		Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var events []persistence.Event
	if err := cursor.All(context.Background(), &events); err != nil {
		return nil, err
	}

	return events, nil
}
