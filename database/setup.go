package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func ensureCollections(db *mongo.Database) error {
	collections := []string{"users"}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, col := range collections {
		if err := db.CreateCollection(ctx, col); err != nil {
			return err
		}
	}

	return nil
}

func Setup() {
	mongo, err := newMongo()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer mongo.close()

	// TODO: stores here
}
