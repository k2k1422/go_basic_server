package main

import (
	"net/http"
	"server/Helloworld"

	"github.com/gorilla/mux"
)

func main() {

	serverMux := mux.NewRouter()

	helloWorldRouter := serverMux.PathPrefix("/demo").Subrouter()

	Helloworld.Route(helloWorldRouter)

	serverMux.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./Build"))))

	err := http.ListenAndServe(":8081", serverMux)

	if err != nil {
		panic(err)
	}

}
