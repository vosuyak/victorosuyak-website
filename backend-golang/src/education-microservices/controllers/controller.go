package controllers

import (
	"fmt"
	"net/http"
	"education/common"
	"education/data"
	"education/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// CreateEducation - creation of a new Education
func CreateEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateEducation")
	var edu models.Education

	err := json.NewDecoder(r.Body).Decode(&edu)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error with decoding form",
		)
		return
	}

	// Create Form
	collection := data.GetCollection("education")
	repo := &EducationRepository{C: collection}
	edu.ID = primitive.NewObjectID()
	errEdu := repo.Create(edu)
	if errEdu != nil {
		common.DisplayError(w, errEdu, http.StatusInternalServerError,
			"error with inserting into database",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, edu)
}

// GetEducation - creation of a new Education
func GetEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetEducation")
}

// GetAllEducation - creation of a new Education
func GetAllEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllEducation")
}

// UpdateEducation - creation of a new Education
func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateEducation")
}

// DeleteEducation - creation of a new Education
func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteEducation")
}
