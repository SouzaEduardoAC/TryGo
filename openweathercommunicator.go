package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	ApiKey             = "ApiKey"
	OpenWeatherBaseUrl = "http://api.openweathermap.org"
)

func GetCurrentWeatherForCity(cn string) map[string]interface{} {
	hc := http.Client{
		Timeout: time.Second * 3,
	}

	url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", OpenWeatherBaseUrl, cn, ApiKey)

	res, err := hc.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var mr map[string]interface{}
	json.NewDecoder(res.Body).Decode(&mr)

	return mr
}
