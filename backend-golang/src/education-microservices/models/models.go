package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Education struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	School     string             `json:"school" bson:"school"`
	StartDate  string             `json:"start_date" bson:"start_date"`
	EndDate    string             `json:"end_date" bson:"end_date"`
	Degree     string             `json:"degree" bson:"degree"`
	City       string             `json:"city" bson:"city"`
	State      string             `json:"state" bson:"state"`
	Online     bool               `json:"online" bson:"online"`
	Website    string             `json:"website" bson:"website"`
	WebsiteURL string             `json:"website_url" bson:"website_url"`
	Price      int                `json:"price" bson:"price"`
	Type       string             `json:"type" bson:"type"`
	Courses    []Course           `json:"courses" bson:"courses"`
}
type Course struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EducationID          primitive.ObjectID `json:"education_id,omitempty" bson:"education_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Author       string             `json:"author" bson:"author"`
	Description string             `json:"description" bson:"description"`
	Topic       string             `json:"topic" bson:"topic"`
	Language       string             `json:"language" bson:"language"`
	Rating      string             `json:"rating" bson:"rating"`
}
