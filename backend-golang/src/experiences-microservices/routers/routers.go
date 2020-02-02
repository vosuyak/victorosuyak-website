package routers

import (
	"github.com/gorilla/mux"
	"experience/controllers"
)

// ExperienceRoutes - routes
func ExperienceRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/experience", controllers.CreateExperience).Methods("POST")
	r.HandleFunc("/experience/{id}", controllers.GetExperience).Methods("GET")
	r.HandleFunc("/experiences", controllers.GetAllExperience).Methods("GET")
	r.HandleFunc("/experience/{id}", controllers.UpdateExperience).Methods("PUT")
	r.HandleFunc("/experience/{id}", controllers.DeleteExperience).Methods("DELETE")
	return r
}
