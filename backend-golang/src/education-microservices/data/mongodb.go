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
	defer cancel()

	credential := options.Credential{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		AuthSource: os.Getenv("DB_NAME"),
	}
	clientOpts := options.Client().ApplyURI("mongodb://"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
		log.Fatal("mongodb://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"))
	}
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to server", err,credential)
	} else {
		fmt.Println("Connected to MongoDB")
	}
	return client
}
