package main

import (
	"net/http"
	"os"
	_ "user/data"
	"user/routers"
)

func main() {
	r := routers.SetUserRoutes()
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
