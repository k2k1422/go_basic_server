package Auth

import "github.com/gorilla/mux"

func Route(serverMux *mux.Router) {

	serverMux.HandleFunc("/v1/login", login)
	serverMux.HandleFunc("/v1/refreshToken", getAccessToken)

}
