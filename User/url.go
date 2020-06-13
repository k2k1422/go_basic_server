package User

import (
	"github.com/gorilla/mux"
)

func Route(serverMux *mux.Router) {
	/*
		This will bind all the url endpoint with the corresponding handle function present in the current package
	*/
	serverMux.HandleFunc("/v1/create", createUser)
	serverMux.HandleFunc("/v1/getUserList", getUserList)

}
