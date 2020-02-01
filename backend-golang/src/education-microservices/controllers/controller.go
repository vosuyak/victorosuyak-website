package controllers

import (
	"fmt"
	"net/http"
	"education/common"
	"education/data"
	"education/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"

)

var (
	collection = data.GetCollection("education")
	repo = &EducationRepository{C: collection}
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

	// link to Repository,define ID of education, pass struct into Create Method

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
	var edu models.Education
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving id",
		)
		return
	}

	edu, errEdu := repo.Get(_id)
	if errEdu != nil {
		common.DisplayError(w, errEdu, http.StatusInternalServerError,
			"err in retrieving education after decoding",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, edu)
}

// GetAllEducation - creation of a new Education
func GetAllEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllEducation")

	educations, err := repo.GetAll()
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error retrieving all educations",
		)
		return
	}

	common.DisplaySuccess(w, true, http.StatusOK, educations)
}

// UpdateEducation - creation of a new Education
func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateEducation")
	var edu models.Education

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}

	errEdu := json.NewDecoder(r.Body).Decode(&edu)
	if errEdu != nil {
		common.DisplayError(w, errEdu, http.StatusInternalServerError,
			"error in decoding education",
		)
		return
	}


	edu.ID = _id
	result := repo.Update(edu)
	if result != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	// show success message in response
	common.DisplaySuccess(w, nil, http.StatusCreated, &edu)
}

// UpdatePatchEducation - creation of a new Education
func UpdatePatchEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdatePatchEducation")
	var edu models.Education

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}

	errEdu := json.NewDecoder(r.Body).Decode(&edu)
	if errEdu != nil {
		common.DisplayError(w, errEdu, http.StatusInternalServerError,
			"error in decoding education",
		)
		return
	}
	fmt.Println(edu)


	edu.ID = _id
	result := repo.Update(edu)
	if result != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	// show success message in response
	common.DisplaySuccess(w, nil, http.StatusCreated, &edu)
}

// DeleteEducation - creation of a new Education
func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteEducation")
		// collect id from router
		vars := mux.Vars(r)
		_id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"error in retrieving id education path",
			)
			return
		}

	
		errEdu := repo.Delete(_id)
		if errEdu != nil {
			common.DisplayError(w, errEdu, http.StatusInternalServerError,
				"error in deleting id education db",
			)
			return
		}
		// show success message in response
		common.DisplaySuccess(w, true, http.StatusOK, "deleted")
}

// DeleteAllEducation - creation of a new Education
func DeleteAllEducation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteAllEducation")
		// collect id from router
		vars := mux.Vars(r)
		_id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"error in retrieving id education path",
			)
			return
		}

	
		errEdu := repo.DeleteAll()
		if errEdu != nil {
			common.DisplayError(w, errEdu, http.StatusInternalServerError,
				"error in deleting id education db",
			)
			return
		}
		// show success message in response
		common.DisplaySuccess(w, true, http.StatusOK, "deleted all")
}