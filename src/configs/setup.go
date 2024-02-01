package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func init() {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		// Handle the error appropriately, such as logging or exiting the program
	}
}

func ConnectDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		fmt.Println("MongoDB connection not established.")
		return nil
	}
	collection := DB.Database(EnvMongoDB()).Collection(collectionName)
	return collection
}
