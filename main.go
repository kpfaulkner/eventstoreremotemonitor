package main

import (
	"flag"
	"fmt"
	"github.com/kpfaulkner/eventstoreremotemonitor/models"
	storage "github.com/kpfaulkner/eventstoreremotemonitor/storage"
	"log"
)


func main() {

	fmt.Printf("So it begins......\n\n\n")

	configFile := flag.String("config", "config.json", "configfile")
	flag.Parse()

	config := models.ReadConfig( *configFile )


	storage, err := storage.NewTableStatsStorage( config )
	if err != nil {
		log.Fatalf("Cannot create storage...  kaboom %s\n", err.Error())
	}


	sc, err := NewStatsCollector(config, storage)
	if err != nil {
		log.Fatalf("Cannot start stats collector...  kaboom %s\n", err.Error())
	}


	webServer, err := NewESMWebServer( config, storage )
	if err != nil {
		log.Fatalf("Cannot start server...  kaboom %s\n", err.Error())
	}


	go sc.Collect()
	webServer.Run()

}

