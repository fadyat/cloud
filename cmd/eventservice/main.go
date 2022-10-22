package main

import (
	"github.com/fadyat/cloud/internal/eventservice/rest"
	"github.com/fadyat/cloud/internal/persistence/db"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var cfg struct {
	RestAPIAddress     string          `envconfig:"REST_API_ADDRESS" default:":8080"`
	DatabaseConnection string          `envconfig:"DB_URI"`
	DatabaseType       db.DatabaseType `envconfig:"DB_TYPE"`
}

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Printf("[DEBUG] Error loading .env.local file: %s", err)
	}

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	log.Println("[INFO] Starting the event service...")
	dbLayer, err := db.NewLayer(cfg.DatabaseType, cfg.DatabaseConnection)
	if err != nil {
		log.Fatal("[ERROR] Could not connect to database layer: ", err)
	} else {
		log.Println("[DEBUG] Connected to database layer.")
	}

	log.Fatal(rest.ServeAPI(cfg.RestAPIAddress, dbLayer))
}
