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

// Run..... and do stuff.
func (esm *ESMServer) Run() {

}




