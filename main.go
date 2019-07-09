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

// displayStats just writes to stdout with various statistics about the EV node.
func displayStats( proc models.ProcStats, cpu bool, mem bool, queues bool) {

	fmt.Printf("==========================================================\n")
	if cpu {
		fmt.Printf("Proc CPU: %f\n", proc.Proc.CPU)
		fmt.Printf("Sys  CPU: %f\n", proc.Sys.CPU)
	}

	if mem {
		fmt.Printf("Proc MEM: %d\n", proc.Proc.Mem)
		fmt.Printf("Sys Free MEM: %d\n", proc.Sys.FreeMem)
	}

	// if want queues, loop through all and display details.
	if queues {
    displayQueues( proc.Es.Queue.IndexCommitter)
		displayQueues( proc.Es.Queue.MainQueue)
		displayQueues( proc.Es.Queue.MasterReplicationService)
		displayQueues( proc.Es.Queue.MonitoringQueue)
		displayQueues( proc.Es.Queue.PersistentSubscriptions)
		displayQueues( proc.Es.Queue.ProjectionCore0)
		displayQueues( proc.Es.Queue.ProjectionCore1)
		displayQueues( proc.Es.Queue.ProjectionCore2)
		displayQueues( proc.Es.Queue.ProjectionsMaster)
	}
	fmt.Printf("==========================================================\n")
}

// displayQueues writes out queue name and some interesting stats.
func displayQueues( queue models.ESQueue) {
	fmt.Printf("Queue: %-30s Avg IPS: %d\tLength: %d\tTotal Items Processed %d\n", queue.QueueName, queue.AvgItemsPerSecond, queue.Length, queue.TotalItemsProcessed)
}
