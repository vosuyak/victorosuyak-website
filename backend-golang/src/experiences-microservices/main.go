package main

import (
	_ "experience/data"
	"os"
	"experience/routers"
	"net/http"
	// "github.com/gorilla/handlers"
)

func main() {
	r := routers.ExperienceRoutes()
	// rather than haveing lochalhost you can also just to "*", however, this is typically for a public api. 
	// If you have a endpoint that uses a cookie, * is not acceptable
	// ch := handlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:4200"})
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
