package mongo

import (
	"context"
	"github.com/fadyat/cloud/internal/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// dbName is the name of the database
	dbName = "db"

	// eventsCollection is the name of the events collection
	eventsCollection = "events"
)

type DBLayer struct {
	client *mongo.Client
}

func (db *DBLayer) Close() error {
	return db.client.Disconnect(context.Background())
}

func NewLayer(connection string) (*DBLayer, error) {
	var clientOptions = options.Client().ApplyURI(connection)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &DBLayer{
		client: client,
	}, nil
}

func (db *DBLayer) CreateEvent(event *persistence.Event) ([]byte, error) {
	s, err := db.client.StartSession()
	defer s.EndSession(context.Background())
	if err != nil {
		return nil, err
	}

	result, err := db.client.Database(dbName).
		Collection(eventsCollection).
		InsertOne(context.Background(), event)

	if err != nil {
		return nil, err
	}

	eventID := result.InsertedID.(primitive.ObjectID)
	return eventID[:], nil
}

func (db *DBLayer) FindEvent(id string) (*persistence.Event, error) {
	eventID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	s, err := db.client.StartSession()
	defer s.EndSession(context.Background())
	if err != nil {
		return nil, err
	}

	var event persistence.Event
	err = db.client.Database(dbName).
		Collection(eventsCollection).
		FindOne(context.Background(), bson.M{"_id": eventID}).
		Decode(&event)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (db *DBLayer) FindEventByName(name string) (*persistence.Event, error) {
	s, err := db.client.StartSession()
	defer s.EndSession(context.Background())
	if err != nil {
		return nil, err
	}

	var event persistence.Event
	err = db.client.Database(dbName).
		Collection(eventsCollection).
		FindOne(context.Background(), bson.M{"name": name}).
		Decode(&event)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (db *DBLayer) FindAll() ([]persistence.Event, error) {
	s, err := db.client.StartSession()
	defer s.EndSession(context.Background())
	if err != nil {
		return nil, err
	}

	cursor, err := db.client.Database(dbName).
		Collection(eventsCollection).
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
