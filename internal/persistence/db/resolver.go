package db

import (
	"fmt"
	"github.com/fadyat/cloud/internal/persistence"
	"github.com/fadyat/cloud/internal/persistence/mongo"
)

type DatabaseType string

func (dbt *DatabaseType) Decode(value string) error {
	*dbt = DatabaseType(value)
	return nil
}

const (
	MongoDB  DatabaseType = "mongodb"
	DynamoDB DatabaseType = "dynamodb"
)

func NewLayer(database DatabaseType, connection string) (persistence.DatabaseHandler, error) {
	switch database {
	case MongoDB:
		return mongo.NewLayer(connection)
	case DynamoDB:
		panic("DynamoDB is not yet implemented")
	default:
		return nil, fmt.Errorf("database type not supported: %s", database)
	}
}
