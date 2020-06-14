package Todo

import "github.com/gorilla/mux"

func Route(serverMux *mux.Router) {

	serverMux.HandleFunc("/v1/getTodoList", getTodoList)
	serverMux.HandleFunc("/v1/addTodo", addTodo)
	serverMux.HandleFunc("/v1/deleteTodo", deleteTodo)

}
