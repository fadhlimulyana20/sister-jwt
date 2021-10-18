package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var db *mongo.Database
var e error

func Connect() error {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://root:password@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db = client.Database("sister_jwt")
	return nil
}

func DbManager() *mongo.Database {
	return db
}
