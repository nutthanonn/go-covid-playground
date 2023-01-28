package datastore

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongodb

func Connect(ctx context.Context) (*mongo.Client, error) {
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@database:27017"))

	if err != nil {
		return nil, err
	}

	return client, nil
}
