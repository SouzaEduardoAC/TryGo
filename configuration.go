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
	f, err := os.Open(fmt.Sprintf("config/config.%s.yml", env))
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
	return cfg
}

func processError(e interface{}) {
	log.Fatal(e)
}
