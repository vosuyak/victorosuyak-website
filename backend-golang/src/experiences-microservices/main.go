package main

import (
	_ "experience/data"
	"os"
	"experience/routers"
	"net/http"
)

func main() {
	r := routers.ExperienceRoutes()
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
