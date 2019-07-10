package main

import (
	"github.com/kpfaulkner/eventstoreremotemonitor/models"
	"net/http"
)

// ESMWebServer will respond to requests (probably from Grafana) about statistics it holds
// about an EventStore cluster.
type ESMWebServer struct {
	config models.Config
	storage StatsStorage

}

// NewESMWebServer create new ESMWebServer... has reference to stats collector.
func NewESMWebServer( config models.Config, storage StatsStorage) (*ESMWebServer, error) {
	esm := ESMWebServer{}

	esm.config = config
	esm.storage = storage

	return &esm, nil
}


// queryES queries EventStore
func (esm *ESMWebServer) queryES() error {

	return nil
}



// CacheStatistics returns a HandlerFunc that gets stats for cache.
func (esm *ESMWebServer) getStats() http.HandlerFunc {


	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// get a payload p := Payload{d}
	//	json.NewEncoder(w).Encode(stats)
	}
}


func (esm *ESMWebServer) setupRoutes() error {

	http.HandleFunc("/stats", esm.getStats())
	return nil
}

// Run starts the webserver that will respond with stats about the EventStore cluster its monitoring.
// Should be a simple webserver.
func (esm *ESMWebServer) Run() {
	esm.setupRoutes()
	http.ListenAndServe(":8080", nil)
}




