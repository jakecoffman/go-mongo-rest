package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Client *mongo.Client
var Database *mongo.Database
var Authors *mongo.Collection

const (
	Name             = "library"
	CollectionAuthor = "author"
)

func Connect(dbName string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	Database = Client.Database(dbName)
	Authors = Database.Collection(CollectionAuthor)
	return nil
}
