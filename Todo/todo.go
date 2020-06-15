package Todo

import (
	"encoding/json"
	"net/http"
	"server/DataAccess"
	"server/DataModels"
	"server/Logging"
	"server/Response"
	"server/SMTP"

	"github.com/google/uuid"
)

func getTodoList(w http.ResponseWriter, r *http.Request) {
	Response.Success(w, r, "204", DataAccess.GetTodoList())
}

func addTodo(w http.ResponseWriter, r *http.Request) {

	var addTodoData DataModels.Todo

	if err := json.NewDecoder(r.Body).Decode(&addTodoData); err != nil {
		Logging.ERROR.Println("Could not read the request body. ", err)
		Response.BadRequest(w, r, "100")
	} else {
		if err := DataAccess.Validator.Struct(addTodoData); err != nil {
			Logging.ERROR.Println("Required data fields in the request body is missing ", err)
			Response.BadRequest(w, r, "101")
		} else {
			addTodoData.ID = uuid.New().String()

			if DataAccess.AddTodo(addTodoData) {
				Logging.INFO.Println("Successfully created the todo")
				Response.Created(w, r, "801")
			} else {
				Logging.ERROR.Println("failed to create the todo")
				Response.InternalServerError(w, r, "802")
			}
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

	var deleteTodoData DataModels.TodoById

	if err := json.NewDecoder(r.Body).Decode(&deleteTodoData); err != nil {
		Logging.ERROR.Println("Could not read the request body. ", err)
		Response.BadRequest(w, r, "100")
	} else {
		if err := DataAccess.Validator.Struct(deleteTodoData); err != nil {
			Logging.ERROR.Println("Required data fields in the request body is missing ", err)
			Response.BadRequest(w, r, "101")
		} else {
			if DataAccess.TodoExistsByID(deleteTodoData.ID) {
				if DataAccess.DeleteTodoByID(deleteTodoData.ID) {
					Logging.INFO.Println("Successfully deleted the todo")
					Response.Created(w, r, "803")
				} else {
					Logging.ERROR.Println("failed to delete the todo")
					Response.InternalServerError(w, r, "804")
				}
			} else {
				Logging.ERROR.Println("failed to delete the todo")
				Response.InternalServerError(w, r, "804")
			}
		}
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) {

	var emailId []string

	if err := json.NewDecoder(r.Body).Decode(&emailId); err != nil {
		Logging.ERROR.Println("Could not read the request body. ", err)
		Response.BadRequest(w, r, "100")
	} else {
		if SMTP.Send(emailId, "(Todo Assiged)",
			SMTP.TodoSMTPTemplate(DataAccess.GetTodoList()),
		) {
			Logging.INFO.Println("Successfully sent email")
			Response.Success(w, r, "805", "")
		} else {
			Logging.ERROR.Println("Failed to sent email")
			Response.InternalServerError(w, r, "806")
		}

	}
}
