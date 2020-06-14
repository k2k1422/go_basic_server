package DataModels

type Todo struct {
	ID       string `json"id"	grom:"column:id;	type:varchar(64);"`
	Task     string `json:"task"	validate:"required"	grom:"column:task;	type:varchar(200);"`
	TaskDate int64  `json:"task_date"	validate:"required"	grom:"column:task_date;	type:int(64);"`
}

type TodoById struct {
	ID       string `json"id"	validate:"required"`
	Task     string `json:"task"`
	TaskDate int64  `json:"task_date"`
}
