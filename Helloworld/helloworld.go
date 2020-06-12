package Helloworld

import (
	"net/http"
	"server/Response"
)

func helloworldDemo(w http.ResponseWriter, r *http.Request) {
	obj := make(map[string]interface{})
	obj["message"] = "Hello world"
	Response.Success(w, r, "700", obj)
}
