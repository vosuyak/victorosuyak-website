package routers

import (
	"github.com/gorilla/mux"
	"skill/controllers"
)

// SkillRoutes - routes
func SkillRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/skill", controllers.CreateSkill).Methods("POST")
	r.HandleFunc("/skill/{id}", controllers.GetSkill).Methods("GET")
	r.HandleFunc("/skills", controllers.GetAllSkill).Methods("GET")
	r.HandleFunc("/skill/{id}", controllers.UpdateSkill).Methods("PUT")
	r.HandleFunc("/skill/{id}", controllers.DeleteSkill).Methods("DELETE")
	return r
}
