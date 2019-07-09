package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	RefreshInSeconds int    `json:"refreshInSeconds"`
	Servers          []struct {
		Name string `json:"name"`
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"servers"`
}

// readConfig reads a JSON file and returns the appropriate config.
func ReadConfig( filePath string ) Config {

	data, err := ioutil.ReadFile( filePath)
	if err != nil {
		log.Fatal(err)
	}
	config := Config{}
	json.Unmarshal(data, &config)

	return config
}
