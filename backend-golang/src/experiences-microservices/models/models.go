package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type     string             `json:"type" bson:"type"`
	TypeTwo    string             `json:"type_two" bson:"type_two"`
	Employer     string             `json:"employer" bson:"employer"`
	EmployerURL     string             `json:"employer_url" bson:"employer_url"`
	StartDate     string             `json:"start_date" bson:"start_date"`
	EndDate     string             `json:"end_date" bson:"end_date"`
	Position     string             `json:"position" bson:"position"`
	TeamSize     int             `json:"team_size" bson:"team_size"`
	Description     string             `json:"description" bson:"description"`
}

