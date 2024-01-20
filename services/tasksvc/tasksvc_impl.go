package tasksvc

import (
	"context"
	"taskreminder/models"
	"time"
)

func (ds *TaskService) HealthCheck() (*models.Health, error) {
	return ds.getDBhealth()
}

func (ds *TaskService) CreateTask(task *models.Task) (*models.TaskResp, error) {
	createdat := time.Now().UnixMilli()
	ctx := context.Background()

	return ds.createTaskDb(ctx, task.Title, task.Description, task.Priority, createdat, task.DueDate, PENDING, task.CreatedBy)

}

func (ds *TaskService) GetAllTask(user string) ([]*models.TaskResp, error) {
	ctx := context.Background()
	return ds.getAllTaskDb(ctx, user)
}

func (ds *TaskService) GetPendingTask(user string) ([]*models.TaskResp, error) {
	ctx := context.Background()
	return ds.getPendingTaskDb(ctx, user)
}

func (ds *TaskService) UpdateTask(updateReq *models.UpdateTaskReq) (*models.TaskResp, error) {
	ctx := context.Background()
	status := getTaskStatusInt(updateReq.Status)
	return ds.updateTaskDb(ctx, updateReq.Id, updateReq.Title, updateReq.Description, updateReq.DueDate, updateReq.Priority, status)
}

func (ds *TaskService) DeleteTask(id int) error {
	ctx := context.Background()
	return ds.deleteTaskDb(ctx, id)
}

func getTaskStatusString(status int) string {
	switch status {
	case DONE:
		return "done"
	case PENDING:
		return "pending"
	case EXPIRED:
		return "expired"
	default:
		return "unknown"
	}
}

func getTaskStatusInt(statusString string) int {
	switch statusString {
	case "done":
		return DONE
	case "pending":
		return PENDING
	case "expired":
		return EXPIRED
	default:
		return -1
	}
}
