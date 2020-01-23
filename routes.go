package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCurrentWeatherByCityName(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(request)
	cn := pathParams["cityName"]
	m := GetCurrentWeatherForCity(cn)
	d := CurrentWeatherDataOf(m)
	r := MakeSingleCurrentWeatherJsonResponse(&d)
	responseWriter.Write(r)
}

func GetCurrentWeatherForAllCitiesOfCountry(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(request)
	cn := pathParams["country"]
	cities := GetCurrentWeatherForAllCitiesOfCountryInJson(cn)
	var ds []CurrentWeatherData
	for _, cityWeather := range cities {
		cw := CurrentWeatherDataOf(cityWeather)
		log.Println(cw)
		ds = append(ds, cw)
	}
	jr := MakeMultipleCurrentWeatherJsonResponse(&ds)
	responseWriter.Write(jr)
}

func MakeMultipleCurrentWeatherJsonResponse(d *[]CurrentWeatherData) []byte {
	r, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		return []byte("")
	}
	return r
}

func MakeSingleCurrentWeatherJsonResponse(d *CurrentWeatherData) []byte {
	r, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		return []byte("")
	}
	return r
}
