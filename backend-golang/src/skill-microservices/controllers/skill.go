package controllers

import (
	"net/http"
	"skill/common"
	"skill/data"
	"skill/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
)

var (
	collectionSkill = data.GetCollection("skills")
)

// CreateSkill - creation of a new Skill
func CreateSkill(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill

	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error with decoding skill",
		)
		return
	}

	// link to Repository,define ID of skill, pass struct into Create Method
	skill.ID = primitive.NewObjectID()
	repo := &SkillRepository{C: collectionSkill}
	errSkill := repo.Create(skill)
	if errSkill != nil {
		common.DisplayError(w, errSkill, http.StatusInternalServerError,
			"error with inserting into database",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, skill)
}

// GetSkill - creation of a new Skill
func GetSkill(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving id",
		)
		return
	}
	repo := &SkillRepository{C: collectionSkill}
	skill, errSkill := repo.Get(_id)
	if errSkill != nil {
		common.DisplayError(w, errSkill, http.StatusInternalServerError,
			"err in retrieving skill after decoding",
		)
		return
	}

	common.DisplaySuccess(w, nil, http.StatusOK, skill)
}

// GetAllSkill - creation of a new Skill
func GetAllSkill(w http.ResponseWriter, r *http.Request) {
	repo := &SkillRepository{C: collectionSkill}
	skills, err := repo.GetAll()
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error retrieving all skills",
		)
		return
	}

	common.DisplaySuccess(w, true, http.StatusOK, skills)
}

// UpdateSkill - creation of a new Skill
func UpdateSkill(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}

	errSkill := json.NewDecoder(r.Body).Decode(&skill)
	if errSkill != nil {
		common.DisplayError(w, errSkill, http.StatusInternalServerError,
			"error in decoding skill",
		)
		return
	}


	skill.ID = _id
	repo := &SkillRepository{C: collectionSkill}
	result := repo.Update(skill)
	if result != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	// show success message in response
	common.DisplaySuccess(w, nil, http.StatusCreated, &skill)
}

// DeleteSkill - creation of a new Skill
func DeleteSkill(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"error in retrieving id skill path",
			)
			return
		}

		repo := &SkillRepository{C: collectionSkill}
		errSkill := repo.Delete(_id)
		if errSkill != nil {
			common.DisplayError(w, errSkill, http.StatusInternalServerError,
				"error in deleting id skill db",
			)
			return
		}
		// show success message in response
		common.DisplaySuccess(w, true, http.StatusOK, "deleted")
}

