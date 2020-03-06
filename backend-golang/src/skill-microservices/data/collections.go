package data

import (
	"context"
	"time"

	"skill/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// GetCollection : Goes to the GetClient Client and returns all Collections
func GetCollection(coll string) *mongo.Collection {
	var collection *mongo.Collection
	switch coll {
	case "skills":
		collection = GetClient().Database(os.Getenv("MONGO_DB_NAME")).Collection("skills")
	}
	return collection
}

// CheckIfIDExist - Check if collection entered along with the ID exist
// in the collection, if not return an error
func CheckIfIDExist(coll string, id primitive.ObjectID) error {
	var result error
	switch coll {
	case "skills":
		var skill models.Skill
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := GetCollection("skills")
		filter := bson.D{
			{"_id", id},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&skill)
	}
	return result
}