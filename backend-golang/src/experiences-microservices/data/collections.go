package data

import (
	"context"
	"time"

	"experience/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"experience/core"
)

// GetCollection : Goes to the GetClient Client and returns all Collections
func GetCollection(coll string) *mongo.Collection {
	var collection *mongo.Collection
	switch coll {
	case "experiences":
		collection = GetClient().Database(core.AppConfig.MgDbName).Collection("experiences")
	}
	return collection
}

// CheckIfIDExist - Check if collection entered along with the ID exist
// in the collection, if not return an error
func CheckIfIDExist(coll string, id primitive.ObjectID) error {
	var result error
	switch coll {
	case "experiences":
		var experience models.Experience
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := GetCollection("experiences")
		filter := bson.D{
			{"_id", id},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&experience)
	}
	return result
}