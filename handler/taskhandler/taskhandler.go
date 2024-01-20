package taskhandler

import (
	"taskreminder/services/tasksvc"
	taskSvc "taskreminder/services/tasksvc"
	"taskreminder/utils"

	"github.com/go-chi/chi"
)

type TaskHandler struct {
	tasksvc taskSvc.TaskSvc
	logger  *utils.Logger
}

func (dh *TaskHandler) RegisterRoutes(router chi.Router) {
	router.Get("/health", dh.health)
	router.Post("/create", dh.createTask)
	router.Get("/getAllTasks", dh.getAllTask)
	router.Get("/getPendingTasks", dh.getPendingTask)
	router.Put("/updateTask", dh.updateTask)
	router.Delete("/deleteTask/id={id}", dh.deleteTask)
}

func NewTaskHandler(taskSvc tasksvc.TaskSvc, logger *utils.Logger) *TaskHandler {
	return &TaskHandler{
		tasksvc: taskSvc,
		logger:  logger,
	}
}
