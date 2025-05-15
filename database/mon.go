package database

import (
	"context"
	"log"
	"time"

	"github.com/copl-uk/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "db"

type mongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func newMongo() (*mongoInstance, error) {
	uri := utils.GetEnv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(databaseName)
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	if err := ensureCollections(db); err != nil {
		return nil, err
	}

	return &mongoInstance{
		Client:   client,
		Database: db,
	}, nil
}

func (m *mongoInstance) close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		log.Fatal("Error disconnecting from MongoDB: ", err)
	}
}
