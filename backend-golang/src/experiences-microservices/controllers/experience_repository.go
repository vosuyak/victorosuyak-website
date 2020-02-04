package controllers

import (
	"context"
	"experience/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExperienceRepository - Repo
type ExperienceRepository struct {
	C *mongo.Collection
}

// Create - Create Method
func (r *ExperienceRepository) Create(exp models.Experience) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.C.InsertOne(ctx, exp)
	return err
}

// Get - Get Method
func (r *ExperienceRepository) Get(id primitive.ObjectID) (models.Experience, error) {
	var exp models.Experience
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.FindOne()
	err := r.C.FindOne(ctx, filter, opts).Decode(&exp)
	return exp, err
}

// GetAll - GetAll Method
func (r *ExperienceRepository) GetAll() ([]models.Experience, error) {
	var experiences []models.Experience
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item,err := r.C.Find(ctx,bson.M{})
	var exp models.Experience
	for item.Next(ctx){
		item.Decode(&exp)
		experiences = append(experiences,exp)
	}
	defer item.Close(ctx)
	return experiences, err
}

// Update - Update Method
func (r *ExperienceRepository) Update(exp models.Experience) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id":exp.ID}
	update := bson.M{
		"$set":bson.M{
			"type": exp.Type,
			"type_two": exp.TypeTwo,
			"employer": exp.Employer,
			"employer_url": exp.EmployerURL,
			"start_date": exp.StartDate,
			"end_date": exp.EndDate,
			"postion": exp.Position,
			"team_size": exp.TeamSize,
			"description": exp.Description,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx,filter,update,opts).Decode(&exp)
	return err
}

// Delete - Delete Method
func (r *ExperienceRepository) Delete(id primitive.ObjectID) error {
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
