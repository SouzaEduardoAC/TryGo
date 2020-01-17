package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func testCall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Get called"}`))
}

func getCurrentWeatherByCityName(responseWriter http.ResponseWriter, request *http.Request) {
	pathParams := mux.Vars(request)
	responseWriter.Header().Set("Content-Type", "application/json")

	cityName := pathParams["cityName"]
	OPEN_WEATHER_KEY := "OpenWeatherKey"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", cityName, OPEN_WEATHER_KEY)

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	openWeatherResponse, getError := spaceClient.Do(req)
	if getError != nil {
		log.Fatal(getError)
	}

	openWeatherResponseBody, readError := ioutil.ReadAll(openWeatherResponse.Body)
	if readError != nil {
		log.Fatal(readError)
	}

	responseWriter.Write([]byte(openWeatherResponseBody))
	//`api.openweathermap.org/data/2.5/weather?id="%s"`, cityId
}
