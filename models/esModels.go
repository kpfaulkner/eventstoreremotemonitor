package models

import "time"


type ESWorker struct {
	QueueName                 string      `json:"queueName"`
	GroupName                 string      `json:"groupName"`
	AvgItemsPerSecond         int         `json:"avgItemsPerSecond"`
	AvgProcessingTime         float64     `json:"avgProcessingTime"`
	CurrentIdleTime           string      `json:"currentIdleTime"`
	CurrentItemProcessingTime interface{} `json:"currentItemProcessingTime"`
	IdleTimePercent           float64     `json:"idleTimePercent"`
	Length                    int         `json:"length"`
	LengthCurrentTryPeak      int         `json:"lengthCurrentTryPeak"`
	LengthLifetimePeak        int         `json:"lengthLifetimePeak"`
	TotalItemsProcessed       int         `json:"totalItemsProcessed"`
	InProgressMessage         string      `json:"inProgressMessage"`
	LastProcessedMessage      string      `json:"lastProcessedMessage"`
}

type ESQueue struct {
	QueueName                 string      `json:"queueName"`
	GroupName                 string      `json:"groupName"`
	AvgItemsPerSecond         int         `json:"avgItemsPerSecond"`
	AvgProcessingTime         int         `json:"avgProcessingTime"`
	CurrentIdleTime           string      `json:"currentIdleTime"`
	CurrentItemProcessingTime interface{} `json:"currentItemProcessingTime"`
	IdleTimePercent           float64     `json:"idleTimePercent"`
	Length                    int         `json:"length"`
	LengthCurrentTryPeak      int         `json:"lengthCurrentTryPeak"`
	LengthLifetimePeak        int         `json:"lengthLifetimePeak"`
	TotalItemsProcessed       int         `json:"totalItemsProcessed"`
	InProgressMessage         string      `json:"inProgressMessage"`
	LastProcessedMessage      string      `json:"lastProcessedMessage"`
}


// ProcStats gives stats on eventstore...
type ProcStats struct {
	Proc struct {
		StartTime            time.Time `json:"startTime"`
		ID                   int       `json:"id"`
		Mem                  int       `json:"mem"`
		CPU                  float64   `json:"cpu"`
		CPUScaled            float64   `json:"cpuScaled"`
		ThreadsCount         int       `json:"threadsCount"`
		ContentionsRate      float64   `json:"contentionsRate"`
		ThrownExceptionsRate float64   `json:"thrownExceptionsRate"`
		Gc                   struct {
			AllocationSpeed   int     `json:"allocationSpeed"`
			Gen0ItemsCount    int     `json:"gen0ItemsCount"`
			Gen0Size          int     `json:"gen0Size"`
			Gen1ItemsCount    int     `json:"gen1ItemsCount"`
			Gen1Size          int     `json:"gen1Size"`
			Gen2ItemsCount    int     `json:"gen2ItemsCount"`
			Gen2Size          int     `json:"gen2Size"`
			LargeHeapSize     int     `json:"largeHeapSize"`
			TimeInGc          float64 `json:"timeInGc"`
			TotalBytesInHeaps int     `json:"totalBytesInHeaps"`
		} `json:"gc"`
		DiskIo struct {
			ReadBytes    int64 `json:"readBytes"`
			WrittenBytes int   `json:"writtenBytes"`
			ReadOps      int   `json:"readOps"`
			WriteOps     int   `json:"writeOps"`
		} `json:"diskIo"`
		TCP struct {
			Connections               int    `json:"connections"`
			ReceivingSpeed            int    `json:"receivingSpeed"`
			SendingSpeed              int    `json:"sendingSpeed"`
			InSend                    int    `json:"inSend"`
			MeasureTime               string `json:"measureTime"`
			PendingReceived           int    `json:"pendingReceived"`
			PendingSend               int    `json:"pendingSend"`
			ReceivedBytesSinceLastRun int    `json:"receivedBytesSinceLastRun"`
			ReceivedBytesTotal        int    `json:"receivedBytesTotal"`
			SentBytesSinceLastRun     int    `json:"sentBytesSinceLastRun"`
			SentBytesTotal            int    `json:"sentBytesTotal"`
		} `json:"tcp"`
	} `json:"proc"`
	Sys struct {
		CPU     float64 `json:"cpu"`
		FreeMem int64   `json:"freeMem"`
		Drive   struct {

			// assumption F drive... yeah, not great. Need to make this generic.
			F struct {
				AvailableBytes int64  `json:"availableBytes"`
				TotalBytes     int64  `json:"totalBytes"`
				Usage          string `json:"usage"`
				UsedBytes      int64  `json:"usedBytes"`
			} `json:"f"`
		} `json:"drive"`
	} `json:"sys"`
	Es struct {
		Checksum           int64 `json:"checksum"`
		ChecksumNonFlushed int64 `json:"checksumNonFlushed"`
		Queue              struct {
			IndexCommitter ESQueue `json:"Index Committer"`
			MainQueue ESQueue `json:"MainQueue"`
			MasterReplicationService ESQueue `json:"Master Replication Service"`
			MonitoringQueue ESQueue `json:"MonitoringQueue"`
			PersistentSubscriptions ESQueue `json:"PersistentSubscriptions"`
			ProjectionCore0 ESQueue `json:"Projection Core #0"`
			ProjectionCore1 ESQueue `json:"Projection Core #1"`
			ProjectionCore2 ESQueue`json:"Projection Core #2"`
			ProjectionsMaster ESQueue `json:"Projections Master"`
			StorageChaser ESQueue `json:"Storage Chaser"`
			StorageReaderQueue1 ESQueue `json:"StorageReaderQueue #1"`
			StorageReaderQueue2 ESQueue `json:"StorageReaderQueue #2"`
			StorageReaderQueue3 ESQueue `json:"StorageReaderQueue #3"`
			StorageReaderQueue4 ESQueue `json:"StorageReaderQueue #4"`
			StorageWriterQueue ESQueue `json:"StorageWriterQueue"`
			Subscriptions ESQueue `json:"Subscriptions"`
			Timer struct {
				QueueName                 string      `json:"queueName"`
				GroupName                 string      `json:"groupName"`
				AvgItemsPerSecond         int         `json:"avgItemsPerSecond"`
				AvgProcessingTime         float64     `json:"avgProcessingTime"`
				CurrentIdleTime           string      `json:"currentIdleTime"`
				CurrentItemProcessingTime interface{} `json:"currentItemProcessingTime"`
				IdleTimePercent           float64     `json:"idleTimePercent"`
				Length                    int         `json:"length"`
				LengthCurrentTryPeak      int         `json:"lengthCurrentTryPeak"`
				LengthLifetimePeak        int         `json:"lengthLifetimePeak"`
				TotalItemsProcessed       int         `json:"totalItemsProcessed"`
				InProgressMessage         string      `json:"inProgressMessage"`
				LastProcessedMessage      string      `json:"lastProcessedMessage"`
			} `json:"Timer"`
			Worker1 ESWorker `json:"Worker #1"`
			Worker2 ESWorker `json:"Worker #2"`
			Worker3 ESWorker `json:"Worker #3"`
			Worker4 ESWorker`json:"Worker #4"`
			Worker5 ESWorker`json:"Worker #5"`
		} `json:"queue"`
		Writer struct {
			LastFlushSize       int     `json:"lastFlushSize"`
			LastFlushDelayMs    float64 `json:"lastFlushDelayMs"`
			MeanFlushSize       int     `json:"meanFlushSize"`
			MeanFlushDelayMs    float64 `json:"meanFlushDelayMs"`
			MaxFlushSize        int     `json:"maxFlushSize"`
			MaxFlushDelayMs     float64 `json:"maxFlushDelayMs"`
			QueuedFlushMessages int     `json:"queuedFlushMessages"`
		} `json:"writer"`
		ReadIndex struct {
			CachedRecord        int `json:"cachedRecord"`
			NotCachedRecord     int `json:"notCachedRecord"`
			CachedStreamInfo    int `json:"cachedStreamInfo"`
			NotCachedStreamInfo int `json:"notCachedStreamInfo"`
			CachedTransInfo     int `json:"cachedTransInfo"`
			NotCachedTransInfo  int `json:"notCachedTransInfo"`
		} `json:"readIndex"`
	} `json:"es"`
	Timestamp time.Time `json:"timestamp"`
}
