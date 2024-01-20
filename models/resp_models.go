package models

type Health struct {
	IsOk     bool `json:"isrunning"`
	DBUpTime int  `json:"dbuptime"`
}

type TaskResp struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	CreatedAt   int64  `json:"createdat"`
	DueDate     int64  `json:"duedate"`
	Status      string `json:"status"`
	CreatedBy   string `json:"createdby"`
}
