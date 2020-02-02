package controllers

import (
	"net/http"
	"education/common"
	"education/data"
	"education/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"

)

var (
	collectionCrs = data.GetCollection("courses")
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
			"error with decoding education_id",
		)
		return
	}

	// link to Repository,define ID of education, pass struct into Create Method
	course.ID = primitive.NewObjectID()
	repo := &CourseRepository{C: collectionCrs}
	errCourse := repo.Create(course)
	if errCourse != nil {
		common.DisplayError(w, errCourse, http.StatusInternalServerError,
			"error with inserting into database",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, course)
}

// GetCourse - creation of a new Course
func GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving id",
		)
		return
	}
	repo := &CourseRepository{C: collectionCrs}
	course, errCrs := repo.Get(_id)
	if errCrs != nil {
		common.DisplayError(w, errCrs, http.StatusInternalServerError,
			"err in retrieving course after decoding",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, course)
}

// GetAllCourses - creation of a new Courses
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	repo := &CourseRepository{C: collectionCrs}
	courses, err := repo.GetAll()
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error retrieving all courses",
		)
		return
	}

	common.DisplaySuccess(w, true, http.StatusOK, courses)
}

// UpdateCourse - creation of a new Course
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}

	errCrs := json.NewDecoder(r.Body).Decode(&course)
	if errCrs != nil {
		common.DisplayError(w, errCrs, http.StatusInternalServerError,
			"error in decoding course",
		)
		return
	}

	// education id require
	errEID := data.CheckIfIDExist("educations",course.EducationID)
	if errEID != nil {
		common.DisplayError(w, errEID, http.StatusInternalServerError,
			"error with decoding education_id",
		)
		return
	}
	
	course.ID = _id
	repo := &CourseRepository{C: collectionCrs}
	result := repo.Update(course)
	if result != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	// show success message in response
	common.DisplaySuccess(w, nil, http.StatusCreated, &course)
}


// DeleteCourse - creation of a new course
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"error in retrieving id course path",
			)
			return
		}

		repo := &CourseRepository{C: collectionCrs}
		errCrs := repo.Delete(_id)
		if errCrs != nil {
			common.DisplayError(w, errCrs, http.StatusInternalServerError,
				"error in deleting id course db",
			)
			return
		}
		// show success message in response
		common.DisplaySuccess(w, true, http.StatusOK, "deleted")
}

