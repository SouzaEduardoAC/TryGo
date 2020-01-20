package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetCurrentWeatherForCity(cn string) map[string]interface{} {
	hc := http.Client{
		Timeout: time.Second * 3,
	}

	c := GetDevelopmentConfiguration()

	url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", c.OpenWeather.Url, cn, c.OpenWeather.Key)
	log.Println(url)
	res, err := hc.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var mr map[string]interface{}
	json.NewDecoder(res.Body).Decode(&mr)

	return mr
}
