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
	r.HandleFunc("/education/{id}/courses", controllers.GetEducationAndCourses).Methods("GET")
	r.HandleFunc("/education/{id}", controllers.UpdateEducation).Methods("PUT")
	r.HandleFunc("/education/{id}", controllers.DeleteEducation).Methods("DELETE")

	r.HandleFunc("/course", controllers.CreateCourse).Methods("POST")
	r.HandleFunc("/courses", controllers.GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", controllers.GetCourse).Methods("GET")
	r.HandleFunc("/course/{id}", controllers.UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", controllers.DeleteCourse).Methods("DELETE")
	return r
}
