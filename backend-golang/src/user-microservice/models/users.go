package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User Struct for Domain Struct
type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrganizationID primitive.ObjectID `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	UserName       string             `json:"user_name" bson:"user_name"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	HashPassword   []byte             `json:"hashpassword,omitempty" bson:"hashpassword,omitempty"`
	FirstName      string             `json:"first_name" bson:"first_name"`
	LastName       string             `json:"last_name"  bson:"last_name"`
	FullName       string             `json:"full_name" bson:"full_name"`
	WorkEmail      string             `json:"work_email" bson:"work_email"`
	WorkPhone      string             `json:"work_phone" bson:"work_phone"`
	PersonalEmail  string             `json:"personal_email" bson:"personal_email"`
	PersonalPhone  string             `json:"personal_phone" bson:"personal_phone"`
	CreatedOn      time.Time          `json:"created_on" bson:"created_on"`
	Country        string             `json:"country" bson:"country"`
	State          string             `json:"state" bson:"state"`
	City           string             `json:"city" bson:"city"`
	Role           string             `json:"role" bson:"role"`
	Domains        []string           `json:"domains" bson:"domains"`
	Services       []string           `json:"services" bson:"services"`
}

type (
	Register struct {
		OrganizationID primitive.ObjectID `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
		UserName       string             `json:"user_name" bson:"user_name"`
		FirstName      string             `json:"first_name" bson:"first_name"`
		LastName       string             `json:"last_name"  bson:"last_name"`
		PersonalEmail  string             `json:"personal_email" bson:"personal_email"`
		WorkEmail      string             `json:"work_email" bson:"work_email"`
		Password       string             `json:"password,omitempty" bson:"password,omitempty"`
		HashPassword   []byte             `json:"hashpassword,omitempty" bson:"hashpassword,omitempty"`
	}

	NewPassword struct {
		UserName    string `json:"user_name" bson:"user_name"`
		Password    string `json:"password,omitempty" bson:"password,omitempty"`
		NewPassword string `json:"new_password,omitempty" bson:"new_password,omitempty"`
	}
	Login struct {
		OrganizationID primitive.ObjectID `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
		WorkEmail      string             `json:"work_email" bson:"work_email"`
		UserName       string             `json:"user_name" bson:"user_name"`
		Password       string             `json:"password,omitempty" bson:"password,omitempty"`
		HashPassword   []byte             `json:"hashpassword,omitempty" bson:"hashpassword,omitempty"`
	}
)
