package main

type CurrentWeatherData struct {
	City_Id             float64
	City_Name           string
	Current_Temperature float64
	Feels_Like          float64
	Country             string
}

func CurrentWeatherDataOf(m map[string]interface{}) CurrentWeatherData {
	c := CurrentWeatherData{
		City_Id:             m["id"].(float64),
		City_Name:           m["name"].(string),
		Current_Temperature: m["main"].(map[string]interface{})["temp"].(float64),
		Feels_Like:          m["main"].(map[string]interface{})["feels_like"].(float64),
		Country:             m["sys"].(map[string]interface{})["country"].(string),
	}
	return c
}

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	OpenWeather struct {
		Key string `yaml:"key"`
		Url string `yaml:"url"`
	} `yaml:"open-weather"`
}
