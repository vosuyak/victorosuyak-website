package data

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection : Goes to the GetClient Client and returns all Collections
func GetCollection(coll string) *mongo.Collection {
	var services *mongo.Collection
	switch coll {
	case "victorosuyak":
		services = GetClient().Database("victorosuyak").Collection("education")
	}
	return services
}
