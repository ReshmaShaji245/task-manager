package tasksvc

import (
	dbservice "taskreminder/database/postgres"
	"taskreminder/models"
	"taskreminder/utils"
)

const (
	PENDING = 1
	DONE    = 2
	EXPIRED = 3
)

type TaskService struct {
	dBSvc  *dbservice.PostgresDBService
	logger *utils.Logger
}

type TaskSvc interface {
	HealthCheck() (*models.Health, error)
	CreateTask(*models.Task) (*models.TaskResp, error)
	GetAllTask(string) ([]*models.TaskResp, error)
	GetPendingTask(string) ([]*models.TaskResp, error)
	UpdateTask(*models.UpdateTaskReq) (*models.TaskResp, error)
	DeleteTask(int) error
}

func NewTaskService(dbservice *dbservice.PostgresDBService, logger *utils.Logger) TaskSvc {
	return &TaskService{
		dBSvc:  dbservice,
		logger: logger,
	}
}
