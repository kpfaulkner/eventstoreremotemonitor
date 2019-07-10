package main

import (
	"github.com/kpfaulkner/eventstoreremotemonitor/models"
	"time"
)


// StatsStorage basically a k/v cache....  until I figure something else out.
type StatsStorage interface {

	// Append stats
	Append( stats models.ProcStats ) error

	// Retrieves the last n seconds from the storage.
	// This means from *now* to *now - n seconds*
	GetLastNSeconds( n int) ([]models.ProcStats, error)
}


// StatsStorage is where all storage happens for the stats.
// both the collector (writer) and web (reader) will have access to this.
// Cant figure out if this needs to be communicated via channels (but cant think of a sensible way of doing that).
// So they (web, collector) will have direct reference to this. Might revisit that later if I figure it out.
type MemStatsStorage struct {
	config models.Config

	// figure out sensible serialisation later... (rolling?)
	statsList []models.ProcStats
}

// NewStatsStorage storage of some sort. Will just be in memory cache for now... need to figure this out more later.
func NewMemStatsStorage( config models.Config ) (*MemStatsStorage, error) {
	ss := MemStatsStorage{}
	ss.config = config
  //ss.cache = make(map[time.Time]models.ProcStats )

  ss.statsList = []models.ProcStats{}

	return &ss, nil
}

// Append append data to time list.
func (mss *MemStatsStorage) Append( stats models.ProcStats ) error {

	// append to end... bet this is inefficnet.
	mss.statsList = append( mss.statsList, stats)
	return nil
}


// GetLastNMinutes...Retrieves the last n seconds from the storage. This means from *now* to *now - n seconds*
// REALLY inefficient... but will do for now.
func (mss *MemStatsStorage) GetLastNSeconds( n int) ([]models.ProcStats, error) {

	now := time.Now().UTC()
	start := now.Add( time.Duration( n * -1) * time.Second)

	index := len(mss.statsList) -1
	found := false
	for {
		element := mss.statsList[index]
		if element.Proc.StartTime.After( start ) {
       found = true

			// go back one.
      index--
		}

		if index < 0 {
			break
		}
  }

	// list we want goes from index to end
	if found {
		return mss.statsList[index+1:], nil
	}

  // nada...  but not an error
  return []models.ProcStats{}, nil
}







