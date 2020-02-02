package data

import (
	"context"
	"time"

	"education/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"education/core"
)

// GetCollection : Goes to the GetClient Client and returns all Collections
func GetCollection(coll string) *mongo.Collection {
	var collection *mongo.Collection
	switch coll {
	case "educations":
		collection = GetClient().Database(core.AppConfig.MgDbName).Collection("education")
	
	case "courses":
		collection = GetClient().Database(core.AppConfig.MgDbName).Collection("courses")
	}
	return collection
}

// CheckIfIDExist - Check if collection entered along with the ID exist
// in the collection, if not return an error
func CheckIfIDExist(coll string, id primitive.ObjectID) error {
	var result error
	switch coll {
	case "educations":
		var education models.Education
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := GetCollection("educations")
		filter := bson.D{
			{"_id", id},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&education)
	case "courses":
		var course models.Course
		ctx, cancel := context.WithTimeout(context.Background(), 05*time.Second)
		defer cancel()
		collection := GetCollection("courses")
		filter := bson.D{
			{"_id", id},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&course)
	}
	return result
}