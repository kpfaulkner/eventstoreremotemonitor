package main

import (
  "github.com/kpfaulkner/eventstoreremotemonitor/models"
	"flag"
	"fmt"
	"log"
)


func main() {

	fmt.Printf("So it begins......\n\n\n")

	configFile := flag.String("config", "", "configfile")
	flag.Parse()

	config := models.ReadConfig( *configFile )

	server, err := NewESMServer( config )
	if err != nil {
		log.Fatalf("Cannot start server...  kaboom %s\n", err.Error())
	}

	server.Run()

}

