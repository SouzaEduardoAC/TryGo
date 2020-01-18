package main

type CurrentWeatherData struct {
	CityId             float64
	CityName           string
	CurrentTemperature float64
	FeelsLike          float64
	Country            string
}

func CurrentWeatherDataOf(m map[string]interface{}) CurrentWeatherData {
	c := CurrentWeatherData{
		CityId:             m["id"].(float64),
		CityName:           m["name"].(string),
		CurrentTemperature: m["main"].(map[string]interface{})["temp"].(float64),
		FeelsLike:          m["main"].(map[string]interface{})["feels_like"].(float64),
		Country:            m["sys"].(map[string]interface{})["country"].(string),
	}
	return c
}
