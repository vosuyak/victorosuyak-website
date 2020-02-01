package main

import (
	_ "education/data"
	"education/routers"
	"net/http"
)

func main() {
	r := routers.EducationRoutes()
	http.ListenAndServe(":8081", r)
}
