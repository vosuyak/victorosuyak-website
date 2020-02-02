package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
	"education/core"
	"os"
)

func init() {
	core.LoadEnv()
	core.LoadAppConfig()
	GetClient()
}

// GetClient - Connect to Client
func GetClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")))
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to server", err)
	} else {
		fmt.Println("Connected to MongoDB")
	}
	return client
}
