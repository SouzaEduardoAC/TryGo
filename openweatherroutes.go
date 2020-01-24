package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterOpenWeatherRouters(router *mux.Router) {
	log.Println("Open Weather API")
	api := router.PathPrefix("/api/open-weather").Subrouter()
	api.HandleFunc("/{cityName}/current-weather", GetCurrentWeatherByCityName).Methods(http.MethodGet)
	api.HandleFunc("/country/{country}/all-cities-weather", GetCurrentWeatherForAllCitiesOfCountry).Methods(http.MethodGet)
	api.HandleFunc("/country/{country}/all-cities-weather-parallel", GetCurrentWeatherForAllCitiesOfCountryParallel).Methods(http.MethodGet)
}
