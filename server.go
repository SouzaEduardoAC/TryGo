package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	configuration := GetDevelopmentConfiguration()
	RegisterOpenWeatherRouters(router)
	log.Fatalln(http.ListenAndServe(":"+configuration.Server.Port, router))
}
