package main

import (
	"github.com/kpfaulkner/eventstoreremotemonitor/models"
	"time"
)


// StatsStorage basically a k/v cache....  until I figure something else out.
type StatsStorage interface {

	// Sets stats in
	Set( stats models.ProcStats ) error

	// Can only get records via timerange, until I figure out if I need something else??
	GetByRange( startTime time.Time, endTime time.Time ) ([]models.ProcStats, error)
}


// StatsStorage is where all storage happens for the stats.
// both the collector (writer) and web (reader) will have access to this.
// Cant figure out if this needs to be communicated via channels (but cant think of a sensible way of doing that).
// So they (web, collector) will have direct reference to this. Might revisit that later if I figure it out.
type MemStatsStorage struct {
	config models.Config
}

// NewStatsStorage storage of some sort. Will just be in memory cache for now... need to figure this out more later.
func NewMemStatsStorage( config models.Config ) (*MemStatsStorage, error) {
	ss := MemStatsStorage{}
	ss.config = config

	return &ss, nil
}


func (mss *MemStatsStorage) Set( stats models.ProcStats ) error {

	return nil
}


func (mss *MemStatsStorage) GetByRange( startTime time.Time, endTime time.Time ) ([]models.ProcStats, error) {

	return nil, nil
}







