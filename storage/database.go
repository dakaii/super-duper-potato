package storage

import (
	"context"
	"graphyy/internal"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var Scooter = "scooter"
var Auth = "auth"

type Storage struct {
	Mongo *mongo.Database
}

// InitDatabase returns a database instance.
func InitDatabase() Storage {
	mongo := initMongo()
	return Storage{
		Mongo: mongo,
	}
}

func initMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(internal.MongoURL))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(internal.DBName)
	collection := db.Collection(Scooter)
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "location", Value: bsonx.String("2dsphere")}},
		},
	}

	// Declare an options object
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err = collection.Indexes().CreateMany(ctx, models, opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
