package main

import (
	_ "education/data"
	"os"
	"education/routers"
	"net/http"
)

func main() {
	r := routers.EducationRoutes()
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
