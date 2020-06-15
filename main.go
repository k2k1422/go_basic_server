package main

import (
	"net/http"
	"server/Auth"
	"server/Helloworld"
	"server/Middleware"
	"server/Todo"
	"server/User"

	"github.com/gorilla/mux"
)

func main() {

	serverMux := mux.NewRouter()

	helloWorldRouter := serverMux.PathPrefix("/demo").Subrouter()
	authRouter := serverMux.PathPrefix("/api/auth").Subrouter()
	userRouter := serverMux.PathPrefix("/api/user").Subrouter()
	todoRouter := serverMux.PathPrefix("/api/todo").Subrouter()

	helloWorldRouter.Use(Middleware.NoAuthLogging)
	authRouter.Use(Middleware.NoAuthLogging)
	userRouter.Use(Middleware.NoAuthLogging)
	todoRouter.Use(Middleware.AuthLogging)

	Helloworld.Route(helloWorldRouter)
	Auth.Route(authRouter)
	User.Route(userRouter)
	Todo.Route(todoRouter)

	serverMux.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./Build"))))

	err := http.ListenAndServe(":8081", serverMux)

	if err != nil {
		panic(err)
	}

}
