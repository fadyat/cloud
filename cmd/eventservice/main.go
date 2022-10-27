package main

import (
	"github.com/fadyat/cloud/internal/eventservice/rest"
	"github.com/fadyat/cloud/internal/persistence/db"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var cfg struct {
	RestAPIAddressHTTP  string          `envconfig:"REST_API_ADDRESS_HTTP" default:":80"`
	RestAPIAddressHTTPS string          `envconfig:"REST_API_ADDRESS_HTTPS" default:":443"`
	DatabaseConnection  string          `envconfig:"DB_URI"`
	DatabaseType        db.DatabaseType `envconfig:"DB_TYPE"`
}

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Printf("[DEBUG] Error loading .env.local file: %s", err)
	}

	if err := envconfig.Process("", &cfg); err != nil {
		log.Panic(err)
	}

	log.Println("[INFO] Starting the event service...")
	dbLayer, err := db.NewLayer(cfg.DatabaseType, cfg.DatabaseConnection)
	defer func() {
		log.Println("[DEBUG] Closing the database connection...")
		if e := dbLayer.Close(); e != nil {
			log.Panic("[ERROR] Error closing the database connection: ", err)
		}
	}()

	if err != nil {
		log.Panic("[ERROR] Could not connect to database layer: ", err)
	} else {
		log.Println("[DEBUG] Connected to database layer.")
	}

	httpErrChan, httpsErrChan := rest.ServeAPI(cfg.RestAPIAddressHTTP, cfg.RestAPIAddressHTTPS, dbLayer)
	defer func() {
		log.Println("[DEBUG] Stopping the http event service...")
		close(httpErrChan)
	}()
	defer func() {
		log.Println("[DEBUG] Stopping the https event service...")
		close(httpsErrChan)
	}()

	select {
	case e := <-httpErrChan:
		log.Panic("[ERROR] HTTP server error: ", e)
	case e := <-httpsErrChan:
		log.Panic("[ERROR] HTTPS server error: ", e)
	}
}
