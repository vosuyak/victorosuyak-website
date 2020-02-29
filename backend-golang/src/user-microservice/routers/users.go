package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"user/controllers"
)

// SetUserRoutes confgs routes for user and returns router
func SetUserRoutes() *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/user/register", controllers.RegisterUser).Methods("POST")
	userRouter.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	userRouter.HandleFunc("/user/profile/{id}", controllers.UserProfile).Methods("GET")
	userRouter.HandleFunc("/user/profile/{id}", controllers.UpdateProfile).Methods("PUT")
	userRouter.HandleFunc("/user/profile/{id}", controllers.DeleteProfile).Methods("DELETE")
	userRouter.PathPrefix("/user").Handler(negroni.New(
		negroni.Wrap(userRouter),
	))
	return userRouter
}
