package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	httpClient = http.Client{
		Timeout: time.Second * 3,
	}
)

func GetCurrentWeatherForCity(cityName string) map[string]interface{} {
	configuration := GetDevelopmentConfiguration()

	url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", configuration.OpenWeather.Url, cityName, configuration.OpenWeather.Key)
	openWeatherResponse, clientError := httpClient.Get(url)
	if clientError != nil {
		log.Fatal(clientError)
	}

	var openWeatherResponseDecoded map[string]interface{}
	json.NewDecoder(openWeatherResponse.Body).Decode(&openWeatherResponseDecoded)

	return openWeatherResponseDecoded
}

func GetCurrentWeatherForAllCitiesOfCountryInJson(countryName string) []map[string]interface{} {
	configuration := GetDevelopmentConfiguration()
	cities := getCitiesFromCountry(countryName)
	citiesToBeUsed := cities[:10]
	response := make([]map[string]interface{}, 0)
	for _, city := range citiesToBeUsed {
		url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", configuration.OpenWeather.Url, city.Name, configuration.OpenWeather.Key)
		openWeatherReponse, _ := httpClient.Get(url)
		var responseDecoded map[string]interface{}
		json.NewDecoder(openWeatherReponse.Body).Decode(&responseDecoded)
		response = append(response, responseDecoded)
	}

	return response
}

func GetCurrentWeatherForCitysInCountry(country string, mainChannel chan CurrentWeatherData) {
	defer close(mainChannel)

	configuration := GetDevelopmentConfiguration()
	cities := getCitiesFromCountry(country)
	citiesToBeUsed := cities[:50]

	var requestChannels = []chan CurrentWeatherData{}

	for i, city := range citiesToBeUsed {
		requestChannels = append(requestChannels, make(chan CurrentWeatherData))
		go GetCurrentWeatherDataForCityRoutine(&configuration, city, requestChannels[i])
	}

	for i := range requestChannels {
		for channel := range requestChannels[i] {
			mainChannel <- channel
		}
	}
}

func GetCurrentWeatherDataForCityRoutine(configuration *Config, city City, openWeatherRequestChannel chan CurrentWeatherData) {
	defer close(openWeatherRequestChannel)
	url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", configuration.OpenWeather.Url, city.Name, configuration.OpenWeather.Key)
	openWeatherResponse, clientError := httpClient.Get(url)
	if clientError != nil {
		log.Fatal(clientError)
	}
	var openWeatherResponseDecoded map[string]interface{}
	json.NewDecoder(openWeatherResponse.Body).Decode(&openWeatherResponseDecoded)
	currentWeatherData := CurrentWeatherDataOf(openWeatherResponseDecoded)
	openWeatherRequestChannel <- currentWeatherData
}

func getCitiesFromCountry(country string) (citiesResponse []City) {
	cities := loadCitiesFromJson()
	for _, city := range cities {
		if city.Country == country {
			citiesResponse = append(citiesResponse, city)
		}
	}

	return
}

func loadCitiesFromJson() (cities []City) {
	file, err := ioutil.ReadFile("./city.list.json")
	if err != nil {
		log.Fatal("File not found")
		return nil
	}

	err = json.Unmarshal([]byte(file), &cities)
	if err != nil {
		log.Fatal("Fail to unmarshall")
		return nil
	}

	return
}
