package data

import (
	"context"
	"time"

	"user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// GetCollection : Goes to the GetClient Client and returns all Collections
func GetCollection(coll string) *mongo.Collection {
	var collection *mongo.Collection
	switch coll {
	case "users":
		collection = GetClient().Database(os.Getenv("MONGO_DB_NAME")).Collection("users")
	}
	return collection
}

// CheckIfIDExist - Check if collection entered along with the ID exist
// in the collection, if not return an error
func CheckIfIDExist(coll string, id primitive.ObjectID) error {
	println("check if exist")

	var result error
	switch coll {
	case "users":
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := GetCollection("users")
		filter := bson.D{
			{"_id", id},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&user)
	}
	return result
}
