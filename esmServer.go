package main

import "github.com/kpfaulkner/eventstoreremotemonitor/models"

type ESMServer struct {

	config models.Config

}

func NewESMServer( config models.Config ) (*ESMServer, error) {
	esm := ESMServer{}

	esm.config = config
	return &esm, nil
}


// queryES queries EventStore
func (esm *ESMServer) queryES() error {

	return nil
}

// Run. Constantly pings ES and gets the stats. Then stores them later for query
// by other systems (probably Grafana).
func (esm *ESMServer) Run() {


}




