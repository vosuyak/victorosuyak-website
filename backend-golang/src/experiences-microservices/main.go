package main

import (
	_ "experience/data"
	"experience/core"
	"experience/routers"
	"net/http"
)

func main() {
	r := routers.ExperienceRoutes()
	http.ListenAndServe(core.AppConfig.Port, r)
}
