package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	c := GetDevelopmentConfiguration()
	RegisterOpenWeatherRouters(r)
	log.Fatalln(http.ListenAndServe(":"+c.Server.Port, r))
}
