package Helloworld

import (
	"net/http"
	"server/Logging"
	"server/Response"
)

func helloworldDemo(w http.ResponseWriter, r *http.Request) {
	obj := make(map[string]interface{})
	obj["message"] = "Hello world"
	Logging.INFO.Println("Succesfully mad a basic hello world with log")
	Response.Success(w, r, "700", obj)
}
