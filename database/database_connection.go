package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// Connect to MongoDB
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		log.Fatal("MONGODB_URL environment variable is not set")
	}
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("grpc_crud").Collection("user")
}
