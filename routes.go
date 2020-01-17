package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	OpenWeatherKey = ""
)

func getCurrentWeatherByCityName(responseWriter http.ResponseWriter, request *http.Request) {
	pathParams := mux.Vars(request)
	responseWriter.Header().Set("Content-Type", "application/json")

	cityName := pathParams["cityName"]
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", cityName, OpenWeatherKey)

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	openWeatherRequest, newRequestError := http.NewRequest(http.MethodGet, url, nil)
	if newRequestError != nil {
		log.Fatal(newRequestError)
	}

	openWeatherResponse, getError := spaceClient.Do(openWeatherRequest)
	if getError != nil {
		log.Fatal(getError)
	}

	var mappedResult map[string]interface{}

	json.NewDecoder(openWeatherResponse.Body).Decode(&mappedResult)

	data := OpenWeatherData{
		CityId:             mappedResult["id"].(float64),
		CityName:           mappedResult["name"].(string),
		CurrentTemperature: mappedResult["main"].(map[string]interface{})["temp"].(float64),
		FeelsLike:          mappedResult["main"].(map[string]interface{})["feels_like"].(float64),
		Country:            mappedResult["sys"].(map[string]interface{})["country"].(string),
	}

	apiResponse, errorResponse := json.Marshal(data)
	if errorResponse != nil {
		log.Fatal(errorResponse)
	}
	responseWriter.Write(apiResponse)
}

type OpenWeatherData struct {
	CityId             float64
	CityName           string
	CurrentTemperature float64
	FeelsLike          float64
	Country            string
}
