package controllers

import (
	"net/http"
	"experience/common"
	"experience/data"
	"experience/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
)

var (
	collectionExp = data.GetCollection("experiences")
)

// CreateExperience - creation of a new Experience
func CreateExperience(w http.ResponseWriter, r *http.Request) {
	var exp models.Experience

	err := json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error with decoding",
		)
		return
	}

	// link to Repository,define ID of experience, pass struct into Create Method
	exp.ID = primitive.NewObjectID()
	repo := &ExperienceRepository{C: collectionExp}
	errExp := repo.Create(exp)
	if errExp != nil {
		common.DisplayError(w, errExp, http.StatusInternalServerError,
			"error with inserting into database",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, exp)
}

// GetExperience - creation of a new Experience
func GetExperience(w http.ResponseWriter, r *http.Request) {
	var exp models.Experience
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving id",
		)
		return
	}
	repo := &ExperienceRepository{C: collectionExp}
	exp, errExp := repo.Get(_id)
	if errExp != nil {
		common.DisplayError(w, errExp, http.StatusInternalServerError,
			"err in retrieving Experience after decoding",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, exp)
}

// GetAllExperience - creation of a new Experience
func GetAllExperience(w http.ResponseWriter, r *http.Request) {
	repo := &ExperienceRepository{C: collectionExp}
	experiences, err := repo.GetAll()
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error retrieving all experiences",
		)
		return
	}

	common.DisplaySuccess(w, true, http.StatusOK, experiences)
}

// UpdateExperience - creation of a new Experience
func UpdateExperience(w http.ResponseWriter, r *http.Request) {
	var exp models.Experience

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}

	errExp := json.NewDecoder(r.Body).Decode(&exp)
	if errExp != nil {
		common.DisplayError(w, errExp, http.StatusInternalServerError,
			"error in decoding experience",
		)
		return
	}


	exp.ID = _id
	repo := &ExperienceRepository{C: collectionExp}
	result := repo.Update(exp)
	if result != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	// show success message in response
	common.DisplaySuccess(w, nil, http.StatusCreated, &exp)
}

// DeleteExperience - creation of a new Experience
func DeleteExperience(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"error in retrieving id Experience path",
			)
			return
		}

		repo := &ExperienceRepository{C: collectionExp}
		errExp := repo.Delete(_id)
		if errExp != nil {
			common.DisplayError(w, errExp, http.StatusInternalServerError,
				"error in deleting id Experience db",
			)
			return
		}
		// show success message in response
		common.DisplaySuccess(w, true, http.StatusOK, "deleted")
}

