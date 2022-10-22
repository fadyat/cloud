package persistence

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Duration  int                `json:"duration"`
	StartDate int64              `json:"start_date"`
	EndDate   int64              `json:"end_date"`
	Location  Location           `json:"location"`
}

type Location struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
