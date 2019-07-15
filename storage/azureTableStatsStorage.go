package storage

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/storage"
	log "github.com/golang/glog"
	"github.com/kpfaulkner/eventstoreremotemonitor/models"
	"io/ioutil"
	"strconv"
)



// TableStatsStorage store against Azure table storage and just serialise the data.
// makes it non-queryable but that's NOT what I'm after.
// Also, given the complexity of the structures, I really don't want to create entities for everything.
type TableStatsStorage struct {
	config models.Config

	tableService storage.TableServiceClient
}

// NewStatsStorage storage of some sort. Will just be in memory cache for now... need to figure this out more later.
func NewTableStatsStorage( config models.Config ) (*TableStatsStorage, error) {
	ss := TableStatsStorage{}
	ss.config = config

	client, err := storage.NewBasicClient("eventstoremonitor", "02NjNytXvKxDnHbbkHtUu3P9sfrBBW/PyI14bJeNxar3DgQiFf2tJ1qhEtNFP1tOkIKLJWbs7wucVDc8pxI9yQ==")

	if err != nil {
		log.Fatal("BOOM!!! ", err)
	}

	tableService := client.GetTableService()
  ss.tableService = tableService

	return &ss, nil
}

func (tss *TableStatsStorage) loadFromDisk() ([]models.ProcStats, error) {

	data, err := ioutil.ReadFile( tss.config.CacheSaveLocation)
	if err != nil {
		log.Errorf("error %s\n",err.Error())
		return []models.ProcStats{}, nil
	}

	procs := []models.ProcStats{}

	json.Unmarshal(data, &procs)

  return procs, nil
}

func (tss *TableStatsStorage) saveToDisk() error {

	/*
	byteArray, err := json.Marshal(mss.statsList)
	if err != nil {
		log.Errorf("Unable to serialise the cache....  kaboom %s\n", err.Error())
		return err
	}


	err = ioutil.WriteFile(mss.config.CacheSaveLocation, byteArray, 0644)
	if err != nil {
		log.Errorf("Unable to save cache to disk....  kaboom %s\n", err.Error())
		return err
	}

*/
  return nil
}

// Append data to table storage
func (tss *TableStatsStorage) Append( stats models.ProcStats ) error {

	byteArray, err := json.Marshal(stats)
	if err != nil {
		log.Errorf("Unable to serialise the cache....  kaboom %s\n", err.Error())
		return err
	}

	table := tss.tableService.GetTableReference("eventstorestats")

	t := stats.Proc.StartTime.UTC().Unix()
	entity := table.GetEntityReference("mypartition",strconv.FormatInt(t, 10))

	props := map[string]interface{}{
		"data":      byteArray,
	}

	entity.Properties = props

	batch := table.NewBatch()

	batch.InsertEntity(entity)
	err = batch.ExecuteBatch()
	if err != nil {
		fmt.Printf("\n\nebatch error %s\n\n", err)

		v, ok := err.(storage.AzureStorageServiceError)
		if ok {
			fmt.Printf("code is %s\n\n", v.Code)
			fmt.Printf("msg is %s\n\n", v.Message)
		}
	}


	// append to end... bet this is inefficnet.
	//mss.statsList = append( tss.statsList, stats)
	return nil
}


// GetLastNMinutes...Retrieves the last n seconds from the storage. This means from *now* to *now - n seconds*
// REALLY inefficient... but will do for now.
func (tss *TableStatsStorage) GetLastNSeconds( n int) ([]models.ProcStats, error) {

	/*
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
*/
  // nada...  but not an error
  return []models.ProcStats{}, nil
}







