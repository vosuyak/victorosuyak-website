package controllers

import (
	"net/http"
	"education/common"
	"education/data"
	"education/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gorilla/mux"

)

var (

)
// CreateCourse - creation of a new Course
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course

	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error with decoding course",
		)
		return
	}

	// education id require
	errEID := data.CheckIfIDExist("educations",course.EducationID)
	if errEID != nil {
		common.DisplayError(w, errEID, http.StatusInternalServerError,
			"error with decoding id",
		)
		return
	}

	// link to Repository,define ID of education, pass struct into Create Method
	course.ID = primitive.NewObjectID()
	collection := data.GetCollection("courses")
	repo := &CourseRepository{C: collection}
	errCourse := repo.Create(course)
	if errCourse != nil {
		common.DisplayError(w, errCourse, http.StatusInternalServerError,
			"error with inserting into database",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, course)
}

// // GetEducation - creation of a new Education
// func GetEducation(w http.ResponseWriter, r *http.Request) {
// 	var edu models.Course
// 	vars := mux.Vars(r)
// 	_id, err := primitive.ObjectIDFromHex(vars["id"])
// 	if err != nil {
// 		common.DisplayError(w, err, http.StatusInternalServerError,
// 			"err in retrieving id",
// 		)
// 		return
// 	}

// 	edu, errEdu := repo.Get(_id)
// 	if errEdu != nil {
// 		common.DisplayError(w, errEdu, http.StatusInternalServerError,
// 			"err in retrieving education after decoding",
// 		)
// 		return
// 	}

// 	common.DisplaySuccess(w, nil, http.StatusOK, edu)
// }

// // GetAllEducation - creation of a new Education
// func GetAllEducation(w http.ResponseWriter, r *http.Request) {
// 	educations, err := repo.GetAll()
// 	if err != nil {
// 		common.DisplayError(w, err, http.StatusInternalServerError,
// 			"error retrieving all educations",
// 		)
// 		return
// 	}

// 	common.DisplaySuccess(w, true, http.StatusOK, educations)
// }

// // UpdateEducation - creation of a new Education
// func UpdateEducation(w http.ResponseWriter, r *http.Request) {
// 	var edu models.Course

// 	vars := mux.Vars(r)
// 	_id, err := primitive.ObjectIDFromHex(vars["id"])
// 	if err != nil {
// 		common.DisplayError(w, err, http.StatusInternalServerError,
// 			"error in retrieving id",
// 		)
// 		return
// 	}

// 	errEdu := json.NewDecoder(r.Body).Decode(&edu)
// 	if errEdu != nil {
// 		common.DisplayError(w, errEdu, http.StatusInternalServerError,
// 			"error in decoding education",
// 		)
// 		return
// 	}


// 	edu.ID = _id
// 	result := repo.Update(edu)
// 	if result != nil {
// 		common.DisplayError(w, err, http.StatusInternalServerError,
// 			"error in updating",
// 		)
// 		return
// 	}
// 	// show success message in response
// 	common.DisplaySuccess(w, nil, http.StatusCreated, &edu)
// }

// // UpdatePatchEducation - creation of a new Education
// func UpdatePatchEducation(w http.ResponseWriter, r *http.Request) {
// 	var edu models.Course

// 	vars := mux.Vars(r)
// 	_id, err := primitive.ObjectIDFromHex(vars["id"])
// 	if err != nil {
// 		common.DisplayError(w, err, http.StatusInternalServerError,
// 			"error in retrieving id",
// 		)
// 		return
// 	}

// 	errEdu := json.NewDecoder(r.Body).Decode(&edu)
// 	if errEdu != nil {
// 		common.DisplayError(w, errEdu, http.StatusInternalServerError,
// 			"error in decoding education",
// 		)
// 		return
// 	}


// 	edu.ID = _id
// 	result := repo.Update(edu)
// 	if result != nil {
// 		common.DisplayError(w, err, http.StatusInternalServerError,
// 			"error in updating",
// 		)
// 		return
// 	}
// 	// show success message in response
// 	common.DisplaySuccess(w, nil, http.StatusCreated, &edu)
// }

// // DeleteEducation - creation of a new Education
// func DeleteEducation(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		_id, err := primitive.ObjectIDFromHex(vars["id"])
// 		if err != nil {
// 			common.DisplayError(w, err, http.StatusInternalServerError,
// 				"error in retrieving id education path",
// 			)
// 			return
// 		}

	
// 		errEdu := repo.Delete(_id)
// 		if errEdu != nil {
// 			common.DisplayError(w, errEdu, http.StatusInternalServerError,
// 				"error in deleting id education db",
// 			)
// 			return
// 		}
// 		// show success message in response
// 		common.DisplaySuccess(w, true, http.StatusOK, "deleted")
// }

