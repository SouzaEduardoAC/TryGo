package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCurrentWeatherByCityName(responseWriter http.ResponseWriter, request *http.Request) {
	pathParams := mux.Vars(request)
	responseWriter.Header().Set("Content-Type", "application/json")
	cn := pathParams["cityName"]
	m := GetCurrentWeatherForCity(cn)
	d := CurrentWeatherDataOf(m)
	r := MakeJsonResponse(&d)
	responseWriter.Write(r)
}

func MakeJsonResponse(d *CurrentWeatherData) []byte {
	r, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		return []byte("")
	}
	return r
}
