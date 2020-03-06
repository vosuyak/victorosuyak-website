package controllers

import (
	"context"
	"fmt"
	"time"

	"user/data"
	"user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	C *mongo.Collection
}


// GetServiceByID - Get service by id method
func (r *UserRepository) GetByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.D{
		{"_id", id},
	}
	opts := options.FindOne().SetSort(bson.D{{"_id", 1}})
	err := r.C.FindOne(ctx, filter, opts).Decode(&user)

	return user, err
}

// Update - Update by ID
func (r *UserRepository) Update(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"first_name":     user.FirstName,
			"last_name":      user.LastName,
			"work_email":     user.WorkEmail,
			"work_phone":     user.WorkPhone,
			"personal_email": user.PersonalEmail,
			"personal_phone": user.PersonalPhone,
			"country":        user.Country,
			"state":          user.State,
			"city":           user.City,
			"role":           user.Role,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx, filter, update, opts).Decode(&user)
	return err
}

// CheckIfUsernameExist :
func (r *UserRepository) CheckIfUsernameExist(username string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.D{
		{"user_name", username},
	}
	opts := options.FindOne().SetSort(bson.D{{"_id", 1}})
	results := r.C.FindOne(ctx, filter, opts).Decode(&user)
	return user, results
}

// UpdatePassword
func (r *UserRepository) UpdatePassword(username string, password *[]byte) error {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Println("update", password)

	filter := bson.D{
		{"user_name", username},
	}
	update := bson.M{
		"$set": bson.M{
			"hashpassword": password,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx, filter, update, opts).Decode(&user)
	return err
}

// FindByUserName - Check if user name exist in users collection
func FindByUserName(coll string, usr string) error {
	var result error
	switch coll {
	case "users":
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 05*time.Second)
		defer cancel()
		collection := data.GetCollection("users")
		filter := bson.D{
			{"user_name", usr},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&user)
	}
	return result
}

// FindByEmail - Check if user name exist in users collection
func FindByEmail(coll string, email string) error {
	var result error
	switch coll {
	case "users":
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 05*time.Second)
		defer cancel()
		collection := data.GetCollection("users")
		filter := bson.D{
			{"personal_email", email},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&user)
	}
	return result
}

// CheckIfEmailExist - Check if collection entered, personal || work with the email exist
// in the collection, if not return an error
func CheckIfEmailExist(use string, email string) error {
	var result error
	var user models.User
	collection := data.GetCollection("users")

	if use == "work" {
		ctx, cancel := context.WithTimeout(context.Background(), 05*time.Second)
		defer cancel()
		filter := bson.D{
			{"work_email", email},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&user)
	} else if use == "personal" {
		ctx, cancel := context.WithTimeout(context.Background(), 05*time.Second)
		defer cancel()
		filter := bson.D{
			{"personal_email", email},
		}
		opts := options.FindOne()
		result = collection.FindOne(ctx, filter, opts).Decode(&user)
	}
	return result
}

// CreateUserAfterRegistration - After registration is successful create user
func CreateUserAfterRegistration(register models.Register, user models.User) (*mongo.InsertOneResult, error) {
	// hash the password entered during registration and remove first password entry
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	register.HashPassword = hashPassword
	// register.Password = ""
	// create user after password has been hashed
	// var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 05*time.Second)
	defer cancel()
	user.ID = primitive.NewObjectID()
	user.OrganizationID = register.OrganizationID
	user.UserName = register.UserName
	user.FirstName = register.FirstName
	user.LastName = register.LastName
	user.FullName = register.FirstName + " " + register.LastName
	user.WorkEmail = register.WorkEmail
	user.PersonalEmail = register.PersonalEmail
	user.HashPassword = hashPassword
	collection := data.GetCollection("users")
	result, err := collection.InsertOne(ctx, user)
	return result, err
}

// Delete 
func (r *UserRepository) Delete(id primitive.ObjectID) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	_, err := r.C.DeleteOne(ctx, filter, opts)
	return err
}