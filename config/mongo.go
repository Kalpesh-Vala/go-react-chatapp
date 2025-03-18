package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// ConnectDB establishes a connection to MongoDB
func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return nil
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("MongoDB ping failed:", err)
		return nil
	}

	fmt.Println("Connected to MongoDB")
	DB = client
	return client
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		DB = ConnectDB()
	}
	return DB.Database("goChatApp").Collection(collectionName)
}

func InitMongo() {
	DB = ConnectDB()
}
