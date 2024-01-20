package taskhandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	handlers "taskreminder/handler"
	models "taskreminder/models"
	"taskreminder/utils"

	"github.com/go-chi/chi"
)

func (dh *TaskHandler) health(w http.ResponseWriter, r *http.Request) {
	response, err := dh.tasksvc.HealthCheck()
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseInternalServerError(w, r, "INTERNAL_SERVER_ERROR", "health check failed", err)
		return
	}

	handlers.APIResponseOK(w, r, response, "health check finished")
}

func (dh *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseBadRequest(w, r, "BAD_REQUEST", "invalid request body", err)
		return
	}
	errstring, err := utils.ValidateStruct(&task)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseUnprocessableEntity(w, r, "FIELDS_MISSING", errstring, err)
	}

	response, err := dh.tasksvc.CreateTask(&task)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseInternalServerError(w, r, "INTERNAL_SERVER_ERROR", "create task failed", err)
		return
	}

	handlers.APIResponseOK(w, r, response, "created task successfully")
}

func (dh *TaskHandler) getAllTask(w http.ResponseWriter, r *http.Request) {
	var user models.CreatedBy
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseBadRequest(w, r, "BAD_REQUEST", "invalid request body", err)
		return
	}
	errstring, err := utils.ValidateStruct(&user)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseUnprocessableEntity(w, r, "FIELDS_MISSING", errstring, err)
		return
	}
	response, err := dh.tasksvc.GetAllTask(user.CreatedBy)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseInternalServerError(w, r, "INTERNAL_SERVER_ERROR", "getting all tasks failed", err)
		return
	}

	handlers.APIResponseOK(w, r, response, "fetched all tasks successfully")
}

func (dh *TaskHandler) getPendingTask(w http.ResponseWriter, r *http.Request) {

	var user models.CreatedBy
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseBadRequest(w, r, "BAD_REQUEST", "invalid request body", err)
		return
	}
	errstring, err := utils.ValidateStruct(&user)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseUnprocessableEntity(w, r, "FIELDS_MISSING", errstring, err)
		return
	}

	response, err := dh.tasksvc.GetPendingTask(user.CreatedBy)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseInternalServerError(w, r, "INTERNAL_SERVER_ERROR", "getting pending tasks failed", err)
		return
	}

	handlers.APIResponseOK(w, r, response, "fetched all pending tasks successfully")
}

func (dh *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {

	var updateReq models.UpdateTaskReq
	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseBadRequest(w, r, "BAD_REQUEST", "invalid request body", err)
		return
	}
	errstring, err := utils.ValidateStruct(&updateReq)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseUnprocessableEntity(w, r, "FIELDS_MISSING", errstring, err)
		return
	}

	response, err := dh.tasksvc.UpdateTask(&updateReq)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseInternalServerError(w, r, "INTERNAL_SERVER_ERROR", "updating task failed", err)
		return
	}

	handlers.APIResponseOK(w, r, response, "updated task successfully")
}

func (dh *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	id = strings.TrimPrefix(id, "id=")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		dh.logger.Error.Print(err)
		handlers.APIResponseBadRequest(w, r, "BAD_REQUEST", "invalid request body", err)
		return
	}

	err1 := dh.tasksvc.DeleteTask(idInt)
	if err1 != nil {
		dh.logger.Error.Print(err1)
		handlers.APIResponseInternalServerError(w, r, "INTERNAL_SERVER_ERROR", "delete task failed", err1)
		return
	}

	handlers.APIResponseOK(w, r, struct{}{}, "deleted task successfully")
}
