package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterOpenWeatherRouters(r *mux.Router) {
	log.Println("Open Weather API")
	api := r.PathPrefix("/api/open-weather").Subrouter()
	api.HandleFunc("/{cityName}/current-weather", GetCurrentWeatherByCityName).Methods(http.MethodGet)
}
