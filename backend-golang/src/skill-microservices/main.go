package main

import (
	"net/http"
	"os"
	_ "skill/data"
	"skill/routers"
)

func main() {
	r := routers.SkillRoutes()
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
