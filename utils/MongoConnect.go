package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseSpecification struct {
	Name       string `bson:"name"`
	SizeOnDisk int64  `bson:"sizeOnDisk"`
	Empty      bool   `bson:"empty"`
}

func MongoConnect() *mongo.Client {
	// Create a context with a timeout of 10 seconds for the MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve MongoDB URI from environment variables
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/nerdingdout").SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		os.Exit(1)
	}

	// Ping the MongoDB server to ensure connectivity
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to ping MongoDB: %v\n", err)
		os.Exit(1)
	}

	return client
}
