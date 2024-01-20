package taskchecksvc

import (
	"context"
	"time"
)

func (ds *TaskCheckSvc) taskCheck() {

	for {
		select {
		case <-time.After(60 * time.Minute):
			ds.updateTaskStatus()
		}
	}
}

func (ds *TaskCheckSvc) updateTaskStatus() {
	currentTime := time.Now()

	conn, err := ds.dbService.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
	}

	sqlQuery := `
			UPDATE Task SET status = ? WHERE duedate < ? AND status = ?
		`
	_, err = conn.Exec(context.Background(), sqlQuery, EXPIRED, currentTime)
	if err != nil {
		ds.logger.Error.Print(err)
	}
}
