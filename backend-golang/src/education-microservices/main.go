package main

import (
	_ "education/data"
	"education/core"
	"education/routers"
	"net/http"
)

func main() {
	r := routers.EducationRoutes()
	http.ListenAndServe(core.AppConfig.Port, r)
}
