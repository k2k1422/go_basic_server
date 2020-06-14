package DataAccess

import (
	"server/DataModels"
	"server/Logging"
)

func GetTodoList() []DataModels.Todo {
	var queryResult []DataModels.Todo
	if err := Connection.Table("sunpharma.TODO").Select("id, task, task_date").Find(&queryResult).Error; err != nil {
		Logging.ERROR.Println("Error while fetching todo data list")
	}
	return queryResult
}

func AddTodo(obj DataModels.Todo) bool {
	if err := Connection.Table("sunpharma.TODO").Create(&obj).Error; err != nil {
		Logging.ERROR.Println("Failed to create the todo")
		return false
	} else {
		Logging.INFO.Println("Successfully created the todo")
		return true
	}
}

func TodoExistsByID(id string) bool {
	/*
		"UserExistsByid" will be used to check whether a user exists in the database or not
	*/
	return !Connection.Table("sunpharma.TODO").Where("id = ?", id).Find(&DataModels.Todo{}).RecordNotFound()
}

func DeleteTodoByID(id string) bool {
	/*
		This will be used to get the user details by user id id
	*/
	if err := Connection.Table("sunpharma.TODO").Where("id = ?", id).Delete(&DataModels.Todo{}).Error; err != nil {
		Logging.ERROR.Println("Error while deleting todo based on id")
		return false
	}
	return true
}
