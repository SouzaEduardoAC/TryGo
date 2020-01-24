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
	cityName := pathParams["cityName"]
	weatherMap := GetCurrentWeatherForCity(cityName)
	currentWeatherData := CurrentWeatherDataOf(weatherMap)
	response := MakeSingleCurrentWeatherJsonResponse(&currentWeatherData)
	responseWriter.Write(response)
}

func GetCurrentWeatherForAllCitiesOfCountry(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(request)
	country := pathParams["country"]
	cities := GetCurrentWeatherForAllCitiesOfCountryInJson(country)
	var currentWeatherDataArray []CurrentWeatherData
	for _, cityWeather := range cities {
		currentWeatherData := CurrentWeatherDataOf(cityWeather)
		currentWeatherDataArray = append(currentWeatherDataArray, currentWeatherData)
	}
	response := MakeMultipleCurrentWeatherJsonResponse(&currentWeatherDataArray)
	responseWriter.Write(response)
}

func GetCurrentWeatherForAllCitiesOfCountryParallel(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(request)
	country := pathParams["country"]
	mainChannel := make(chan CurrentWeatherData)
	go GetCurrentWeatherForCitysInCountry(country, mainChannel)

	var channelData []CurrentWeatherData
	for element := range mainChannel {
		channelData = append(channelData, element)
	}

	response := MakeMultipleCurrentWeatherJsonResponse(&channelData)
	responseWriter.Write(response)
}

func MakeMultipleCurrentWeatherJsonResponse(data *[]CurrentWeatherData) []byte {
	response, error := json.Marshal(data)
	if error != nil {
		log.Fatal(error)
		return []byte("")
	}
	return response
}

func MakeSingleCurrentWeatherJsonResponse(data *CurrentWeatherData) []byte {
	response, error := json.Marshal(data)
	if error != nil {
		log.Fatal(error)
		return []byte("")
	}
	return response
}
