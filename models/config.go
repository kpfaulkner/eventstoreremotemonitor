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
	CacheSaveLocation string `json:"cacheSaveLocation"`

	// DB for storing stats results.
	DBConnectionString string `json:"dbConnectionString"`

	// has list of servers... but really should only require 1.
	// Will contact the gossip endpoint and get the data from the server.
	// Will only use other entries here IF the first one appears dead.
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
