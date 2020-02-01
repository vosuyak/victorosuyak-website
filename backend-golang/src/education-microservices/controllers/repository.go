package controllers

import (
	"context"
	"education/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EducationRepository - Repo
type EducationRepository struct {
	C *mongo.Collection
}

// Create - Create Method
func (r *EducationRepository) Create(edu models.Education) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.C.InsertOne(ctx, edu)
	return err
}

// Get - Get Method
func (r *EducationRepository) Get(id primitive.ObjectID) (models.Education, error) {
	var edu models.Education
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.FindOne()
	err := r.C.FindOne(ctx, filter, opts).Decode(&edu)
	return edu, err
}

// GetAll - GetAll Method
func (r *EducationRepository) GetAll() ([]models.Education, error) {
	var educations []models.Education
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item,err := r.C.Find(ctx,bson.M{})
	var edu models.Education
	for item.Next(ctx){
		item.Decode(&edu)
		educations = append(educations,edu)
	}
	defer item.Close(ctx)
	return educations, err
}

// Update - Update Method
func (r *EducationRepository) Update(edu models.Education) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id":edu.ID}
	update := bson.M{
		"$set":bson.M{
			"school": edu.School,
			"start_date": edu.StartDate,
			"end_date": edu.EndDate,
			"degree": edu.Degree,
			"city": edu.City,
			"state": edu.State,
			"online": edu.Online,
			"website": edu.Website,
			"website_url": edu.WebsiteURL,
			"price": edu.Price,
			"type": edu.Type,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx,filter,update,opts).Decode(&edu)
	return err
}

// Patch - Patch Method
func (r *EducationRepository) Patch(edu models.Education, key string, val interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": edu.ID}
	update := bson.M{
		"$set":bson.M{
			key : val,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx,filter,update,opts).Decode(&edu)
	return err
}

// Delete - Delete Method
func (r *EducationRepository) Delete(id primitive.ObjectID) error {
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
