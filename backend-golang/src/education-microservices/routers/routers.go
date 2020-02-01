package routers

import (
	"github.com/gorilla/mux"
	"education/controllers"
)

// EducationRoutes - routes
func EducationRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/education", controllers.CreateEducation).Methods("POST")
	r.HandleFunc("/education/{id}", controllers.GetEducation).Methods("GET")
	r.HandleFunc("/educations", controllers.GetAllEducation).Methods("GET")
	r.HandleFunc("/education/{id}", controllers.UpdateEducation).Methods("PUT")
	r.HandleFunc("/education/{id}", controllers.UpdatePatchEducation).Methods("PATCH")
	r.HandleFunc("/education/{id}", controllers.DeleteEducation).Methods("DELETE")
	r.HandleFunc("/education/course", controllers.CreateCourse).Methods("POST")
	// r.HandleFunc("/education/{id}", controllers.GetCourse).Methods("GET")
	// r.HandleFunc("/educations", controllers.GetAllEducation).Methods("GET")
	// r.HandleFunc("/education/{id}", controllers.UpdateEducation).Methods("PUT")
	// r.HandleFunc("/education/{id}", controllers.UpdatePatchEducation).Methods("PATCH")
	// r.HandleFunc("/education/{id}", controllers.DeleteEducation).Methods("DELETE")
	return r
}
