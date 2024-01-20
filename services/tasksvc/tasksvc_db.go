package tasksvc

import (
	"context"
	"log"
	"taskreminder/models"
)

func (ds *TaskService) getDBhealth() (*models.Health, error) {
	conn, err := ds.dBSvc.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	query := `SELECT checkpoints_timed AS uptime FROM pg_stat_bgwriter;
	`
	log.Println("conn: ", conn)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	var (
		uptime int
	)

	for rows.Next() {
		err := rows.Scan(&uptime)
		if err != nil {
			ds.logger.Error.Print(err)
			return nil, err
		}
	}

	return &models.Health{
		IsOk:     true,
		DBUpTime: uptime,
	}, nil
}

func (ds *TaskService) createTaskDb(ctx context.Context, title, description string, priority int, createdat, duedate int64, status int, createdby string) (*models.TaskResp, error) {

	conn, err := ds.dBSvc.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	sqlQuery := `
		INSERT INTO Task (title, description, priority, createdat, duedate, status, createdby)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, title, description, priority, createdat, duedate, status, createdBy
	`
	row := conn.QueryRow(ctx, sqlQuery,
		title, description, priority, createdat, duedate, status, createdby)

	createdTask := &models.TaskResp{}
	var retrievedStatus int
	err1 := row.Scan(
		&createdTask.Id,
		&createdTask.Title,
		&createdTask.Description,
		&createdTask.Priority,
		&createdTask.CreatedAt,
		&createdTask.DueDate,
		&retrievedStatus,
		&createdTask.CreatedBy,
	)
	createdTask.Status = getTaskStatusString(retrievedStatus)

	if err1 != nil {
		return nil, err1
	}

	return createdTask, nil
}

func (ds *TaskService) getPendingTaskDb(ctx context.Context, user string) ([]*models.TaskResp, error) {

	conn, err := ds.dBSvc.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	sqlQuery := `
		SELECT id, title, description, priority, createdat, duedate, createdby
		FROM Task WHERE status=$1 AND createdby=$2
	`

	rows, err := conn.Query(ctx, sqlQuery, PENDING, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*models.TaskResp{}
	for rows.Next() {
		task := models.TaskResp{}
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.CreatedAt,
			&task.DueDate,
			&task.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		task.Status = "pending"
		tasks = append(tasks, &task)
		log.Println("got row", tasks)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ds *TaskService) getAllTaskDb(ctx context.Context, user string) ([]*models.TaskResp, error) {
	conn, err := ds.dBSvc.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	sqlQuery := `
		SELECT id, title, description, priority, createdat, duedate, status, createdby
		FROM Task WHERE createdby=$1
	`

	rows, err := conn.Query(ctx, sqlQuery, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*models.TaskResp{}
	for rows.Next() {
		task := models.TaskResp{}
		var status int
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.CreatedAt,
			&task.DueDate,
			&status,
			&task.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		task.Status = getTaskStatusString(status)
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ds *TaskService) updateTaskDb(ctx context.Context, id int, title, description string, duedate int64, priority int, status int) (*models.TaskResp, error) {
	conn, err := ds.dBSvc.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	sqlQuery := `
		UPDATE Task SET title=$1, description=$2, duedate=$3, priority=$4, status=$5
		WHERE id=$6 RETURNING id, title, description, priority, createdat, duedate, status, createdby
	`

	row := conn.QueryRow(ctx, sqlQuery, title, description, duedate, priority, status, id)
	if err != nil {
		return nil, err
	}
	task := &models.TaskResp{}
	var retrievedStatus int
	err1 := row.Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.CreatedAt,
		&task.DueDate,
		&retrievedStatus,
		&task.CreatedBy,
	)
	task.Status = getTaskStatusString(retrievedStatus)

	if err1 != nil {
		return nil, err1
	}

	return task, nil
}

func (ds *TaskService) deleteTaskDb(ctx context.Context, id int) error {
	conn, err := ds.dBSvc.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return err
	}
	sqlQuery := `DELETE FROM Task WHERE id = $1`

	_, err1 := conn.Exec(ctx, sqlQuery, id)
	if err1 != nil {
		return err1
	}
	return nil
}
