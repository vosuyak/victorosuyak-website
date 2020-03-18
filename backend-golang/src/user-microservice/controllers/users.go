package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"user/common"
	"user/data"
	"user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser : Create a Service
// URI - /user/register
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// decode incoming request payload
	var register models.Register
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error with decoding user for signing up",
		)
		return
	}

	//Check if username or email provided already exist
	userResult := FindByUserName("users", register.UserName)
	emailResult := CheckIfEmailExist("personal", register.PersonalEmail)
	if userResult == nil || emailResult == nil {
		var message string
		// if username and email already exist
		if userResult == nil && emailResult == nil {
			message = "Username and Email Already exist"
		}
		// if email exist but not username
		if userResult != nil && emailResult == nil {
			message = "Email Already exist"
		}
		// if username exist but not email
		if userResult == nil && emailResult != nil {
			message = "Username Already exist"
		}
		common.DisplayCustomResp(w, message, http.StatusFound, nil)
	} else {
		// check if password is complex

		// if username exist and email exist create user
		result, err := CreateUserAfterRegistration(register, user)
		if err != nil {
			common.DisplayError(w, err, http.StatusInternalServerError,
				"err in creating the user",
			)
			return
		}
		common.DisplaySuccess(w, result, http.StatusCreated, nil)
	}
}
// LoginUser : Read All Services
// URI - /user/login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var login models.Login

	//decode payload
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in decoding login",
		)
		return
	}

	// check if username exist in table
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := data.GetCollection("users")
	filter := bson.M{"user_name": login.UserName}
	opts := options.FindOne()
	_ = collection.FindOne(ctx, filter, opts).Decode(&user)


	// check if username exist`
	userResult := FindByUserName("users", login.UserName)

	// compare current hash to login password
	passErr := bcrypt.CompareHashAndPassword(user.HashPassword, []byte(login.Password))
	if userResult != nil || passErr != nil {
		if userResult != nil && passErr == nil {
			common.DisplayCustomResp(w, "incorrect username", http.StatusInternalServerError, nil)
		}
		if userResult == nil && passErr != nil {
			common.DisplayError(w, passErr, http.StatusInternalServerError,
				"wrong password",
			)
			return
		}
		if userResult != nil && passErr != nil {
			common.DisplayCustomResp(w, "username or password is incorrect", http.StatusInternalServerError, nil)
		}
	}

	//set tokens
	tokenString, err := common.GenerateJWT(user.UserName, user.Role, user.PersonalEmail)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"unable to generate token string",
		)
	}
	//set session
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(45 * time.Minute),
	})

	if userResult == nil && passErr == nil {
		// user.Password = ""
		// user.HashPassword = []byte("")
		common.DisplayCustomResp(w, "logged in", http.StatusOK, user)
	}
}
// UserProfile : Read One Service
// URI - /user/profile
func UserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving id",
		)
		return
	}
	collection := data.GetCollection("users")
	repo := &UserRepository{C: collection}
	user, err := repo.GetByID(_id)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"err in retrieving user after decoding",
		)
		return
	}
	common.DisplaySuccess(w, nil, http.StatusOK, user)
}
// UpdateProfile : Update user information
// URI | PUT- /user/profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}
	errForm := json.NewDecoder(r.Body).Decode(&user)
	if errForm != nil {
		common.DisplayError(w, errForm, http.StatusInternalServerError,
			"error in decoding user",
		)
		return
	}
	user.ID = _id
	collection := data.GetCollection("users")
	repo := &UserRepository{C: collection}
	errUser := repo.Update(user)
	if errUser != nil {
		common.DisplayError(w, errUser, http.StatusInternalServerError,
			"error in updating",
		)
		return
	}
	common.DisplaySuccess(w, nil, http.StatusCreated, user)
}
// DeleteProfile
// URI | Delete /user/profile
func DeleteProfile(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	_id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in retrieving id",
		)
		return
	}
	collection := data.GetCollection("users")
	repo := &UserRepository{C: collection}
	err = repo.Delete(_id)
	if err != nil {
		common.DisplayError(w, err, http.StatusInternalServerError,
			"error in deleting",
		)
		return
	}
	common.DisplaySuccess(w, nil, http.StatusCreated, nil)
}