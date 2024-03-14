package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInfrastructure struct {
	Client *mongo.Client
}

func NewMongoDBInfrastructure() (*MongoDBInfrastructure, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return &MongoDBInfrastructure{
		Client: client,
	}, nil
}

// Disconnect closes the MongoDB connection.
func (mongoDBInfrastructure *MongoDBInfrastructure) Disconnect() {
	if err := mongoDBInfrastructure.Client.Disconnect(context.Background()); err != nil {
		log.Println("Error disconnecting from MongoDB:", err)
	}
}

// GetMongoDBClient returns the MongoDB client instance.
func (mongoDBInfrastructure *MongoDBInfrastructure) GetMongoDBClient() *mongo.Client {
	return mongoDBInfrastructure.Client
}

func (mongoDBInfrastructure *MongoDBInfrastructure) GetContext() context.Context {
	return context.Background()
}
