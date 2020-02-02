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

// CourseRepository - Repo
type CourseRepository struct {
	C *mongo.Collection
}

// Create - Create Method
func (r *CourseRepository) Create(course models.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.C.InsertOne(ctx, course)
	return err
}

// Get - Get Method
func (r *CourseRepository) Get(id primitive.ObjectID) (models.Course, error) {
	var course models.Course
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.FindOne()
	err := r.C.FindOne(ctx, filter, opts).Decode(&course)
	return course, err
}

// GetAll - GetAll Method
func (r *CourseRepository) GetAll() ([]models.Course, error) {
	var courses []models.Course
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item,err := r.C.Find(ctx,bson.M{})
	var course models.Course
	for item.Next(ctx){
		item.Decode(&course)
		courses = append(courses,course)
	}
	defer item.Close(ctx)
	return courses, err
}

// Update - Update Method
func (r *CourseRepository) Update(course models.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id":course.ID}
	update := bson.M{
		"$set":bson.M{
			"title": course.Title,
			"author": course.Author,
			"descriotion": course.Description,
			"topic": course.Topic,
			"language": course.Language,
			"rating": course.Rating,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx,filter,update,opts).Decode(&course)
	return err
}

// Delete - Delete Method
func (r *CourseRepository) Delete(id primitive.ObjectID) error {
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
