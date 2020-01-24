package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func GetDevelopmentConfiguration() Config {
	return getConfiguration("develop")
}

func GetProductionConfiguration() Config {
	return getConfiguration("prod")
}

func getConfiguration(env string) Config {
	file, err := os.Open(fmt.Sprintf("config/config.%s.yml", env))
	if err != nil {
		processError(err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		processError(err)
	}
	return config
}

func processError(e interface{}) {
	log.Fatal(e)
}
