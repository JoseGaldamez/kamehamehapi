package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabaseClient(logger *log.Logger) *mongo.Client {
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_DB_URL")
	if mongoURI == "" {
		logger.Fatal("You must provide a MONGO_DB_URL to connect to the database")
	}

	mongoOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	logger.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
