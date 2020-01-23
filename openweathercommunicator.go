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
	hc = http.Client{
		Timeout: time.Second * 3,
	}
)

func GetCurrentWeatherForCity(cn string) map[string]interface{} {
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

func GetCurrentWeatherForAllCitiesOfCountryInJson(cn string) []map[string]interface{} {
	c := GetDevelopmentConfiguration()
	cities := getCitiesFromCountry(cn)
	citiesToBeUsed := cities[:10]
	r := make([]map[string]interface{}, 0)
	for _, city := range citiesToBeUsed {
		url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", c.OpenWeather.Url, city.Name, c.OpenWeather.Key)
		log.Println(url)
		res, _ := hc.Get(url)
		var rb map[string]interface{}
		json.NewDecoder(res.Body).Decode(&rb)
		r = append(r, rb)
	}

	return r
}

func GetCurrentWeatherForCitysInCountry(cn string, rchan chan CurrentWeatherData) {
	defer close(rchan)

	c := GetDevelopmentConfiguration()
	cities := getCitiesFromCountry(cn)
	citiesToBeUsed := cities[:10]

	var results = []chan CurrentWeatherData{}

	for i, city := range citiesToBeUsed {
		results = append(results, make(chan CurrentWeatherData))
		go GetCurrentWeatherDataForCityRoutine(&c, city, results[i])
	}

	for i := range results {
		for r1 := range results[i] {
			rchan <- r1
		}
	}
}

func GetCurrentWeatherDataForCityRoutine(c *Config, city City, rchan chan CurrentWeatherData) {
	defer close(rchan)
	url := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", c.OpenWeather.Url, city.Name, c.OpenWeather.Key)
	res, err := hc.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var rb map[string]interface{}
	json.NewDecoder(res.Body).Decode(&rb)
	cwd := CurrentWeatherDataOf(rb)
	rchan <- cwd
}

func getCitiesFromCountry(cn string) (c []City) {
	cities := loadCitiesFromJson()
	for _, e := range cities {
		if e.Country == cn {
			c = append(c, e)
		}
	}

	return
}

func loadCitiesFromJson() (c []City) {
	file, err := ioutil.ReadFile("./city.list.json")
	if err != nil {
		log.Fatal("File not found")
		return nil
	}

	err = json.Unmarshal([]byte(file), &c)
	if err != nil {
		log.Fatal("Fail to unmarshall")
		return nil
	}

	return
}
