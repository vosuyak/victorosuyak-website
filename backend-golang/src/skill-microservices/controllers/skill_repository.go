package controllers

import (
	"context"
	"skill/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SkillRepository - Repo
type SkillRepository struct {
	C *mongo.Collection
}

// Create - Create Method
func (r *SkillRepository) Create(skill models.Skill) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.C.InsertOne(ctx, skill)
	return err
}

// Get - Get Method
func (r *SkillRepository) Get(id primitive.ObjectID) (models.Skill, error) {
	var skill models.Skill
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.FindOne()
	err := r.C.FindOne(ctx, filter, opts).Decode(&skill)
	return skill, err
}

// GetAll - GetAll Method
func (r *SkillRepository) GetAll() ([]models.Skill, error) {
	var skills []models.Skill
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item,err := r.C.Find(ctx,bson.M{})
	var skill models.Skill
	for item.Next(ctx){
		item.Decode(&skill)
		skills = append(skills,skill)
	}
	defer item.Close(ctx)
	return skills, err
}

// Update - Update Method
func (r *SkillRepository) Update(skill models.Skill) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id":skill.ID}
	update := bson.M{
		"$set":bson.M{
			"language": skill.Language,
			"image": skill.Image,
			"background_color": skill.BackgroundColor,
			"years_experience": skill.YearsExperience,
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := r.C.FindOneAndUpdate(ctx,filter,update,opts).Decode(&skill)
	return err
}

// Delete - Delete Method
func (r *SkillRepository) Delete(id primitive.ObjectID) error {
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

