package main

import (
	"net/http"
	"education/routers"
	_"education/data"

)

func main() {
	r := routers.EducationRoutes()
	http.ListenAndServe(":8081", r)
}
