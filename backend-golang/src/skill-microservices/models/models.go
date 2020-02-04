package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Skill struct{
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Language string `json:"language" bson:"language"`
	YearsExperience int `json:"years_experience" bson:"years_experience"`
}