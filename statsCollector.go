package main

import (
	"encoding/json"
	"fmt"
	log "github.com/golang/glog"
	"github.com/kpfaulkner/eventstoreremotemonitor/models"
	"io/ioutil"
	"net/http"
	"time"
)

// StatsCollector queries the ES cluster and gets all the juicy goss about whats happening.
type StatsCollector struct {
	config models.Config
  storage StatsStorage
}

func NewStatsCollector( config models.Config, storage StatsStorage ) (*StatsCollector, error) {
	sc := StatsCollector{}

	sc.config = config
	sc.storage = storage
	return &sc, nil
}

// queryEventStore executes a "stats" GET against a ES node
// TODO(kpfaulkner) confirm if it matters which node?
func (sc *StatsCollector) queryEventStore() (*models.ProcStats, error) {
	// just use first server....  figure out more robust logic later. TODO(kpfaulkner)
	url := fmt.Sprintf("http://%s:%d/stats", sc.config.Servers[0].IP, sc.config.Servers[0].Port)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth( sc.config.Username, sc.config.Password)
	resp, err := client.Do(req)
	if err != nil{
		log.Errorf("Unable to download stats %s\n", err.Error())
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	proc := models.ProcStats{}
	json.Unmarshal(bodyBytes, &proc)
  return &proc, nil
}

// Collect will go and start collecting stats from the ES cluster.
// Expecting this to really be run in a goroutine.
func (sc *StatsCollector) Collect() error {

	for {

    stats, err := sc.queryEventStore()
    if err != nil {
    	log.Errorf("Unable to retrieve stats %s\n", err.Error())
    	continue
    }

    // store the stats.... SOMEWHERE....  still need to figure this out.
		err = sc.storage.Append(*stats)
		if err != nil {
			log.Errorf("Unable to store stats %s\n", err.Error())
		}

		time.Sleep( time.Duration(sc.config.RefreshInSeconds) * time.Second)
	}

	return nil
}
