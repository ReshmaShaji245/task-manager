package taskchecksvc

import (
	"sync"
	dbservice "taskreminder/database/postgres"
	"taskreminder/utils"
)

const (
	PENDING = 1
	DONE    = 2
	EXPIRED = 3
)

type TaskCheckSvc struct {
	dbService     *dbservice.PostgresDBService
	logger        *utils.Logger
	activeThreads sync.WaitGroup
}

type TaskCheckSvcInterface interface {
	Start()
	Stop()
}

func NewTaskCheckSvc(dbService *dbservice.PostgresDBService, logger *utils.Logger) TaskCheckSvcInterface {
	return &TaskCheckSvc{
		dbService: dbService,
		logger:    logger,
	}
}

func (ds *TaskCheckSvc) Start() {
	ds.activeThreads.Add(1)
	go func() {
		ds.activeThreads.Wait()
		ds.taskCheck()
	}()
}

func (ds *TaskCheckSvc) Stop() {
	ds.activeThreads.Done()
}
