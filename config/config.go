package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var DBName string

func ConnectDB() {
	_ = godotenv.Load()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set")
	}

	DBName = os.Getenv("DB_NAME")
	if DBName == "" {
		DBName = "taskdb"
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI).
		SetConnectTimeout(20 * time.Second).
		SetServerSelectionTimeout(20 * time.Second)

	// CONNECT (v2 driver)
	var err error
	Client, err = mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// context only for ping
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	log.Printf("✅ Connected to MongoDB! Database: %s", DBName)
}
