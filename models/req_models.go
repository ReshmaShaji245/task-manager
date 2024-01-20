package models

type Task struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Priority    int    `json:"priority" validate:"required"`
	DueDate     int64  `json:"duedate" validate:"required"`
	CreatedBy   string `json:"createdby" validate:"required"`
}

type CreatedBy struct {
	CreatedBy string `json:"createdby" validate:"required"`
}

type UpdateTaskReq struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Priority    int    `json:"priority" validate:"required"`
	DueDate     int64  `json:"duedate" validate:"required"`
	Status      string `json:"status" validate:"required"`
}
