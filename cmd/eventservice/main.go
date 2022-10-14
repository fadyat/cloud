package main

import (
	"github.com/fadyat/cloud/internal/eventservice/rest"
	"github.com/fadyat/cloud/internal/persistence/mongo"
	"log"
)

func main() {
	log.Println("Starting the event service...")
	mongoLayer, err := mongo.NewDBLayer("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Could not connect to database layer: ", err)
	} else {
		log.Println("Connected to database layer.")
	}

	log.Fatal(rest.ServeAPI(":8080", mongoLayer))
}
