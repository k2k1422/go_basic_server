package Helloworld

import (
	"github.com/gorilla/mux"
)

func Route(serverMux *mux.Router) {
	serverMux.HandleFunc("/helloworld", helloworldDemo)

}
