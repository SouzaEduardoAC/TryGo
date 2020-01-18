package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	RegisterOpenWeatherRouters(r)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
