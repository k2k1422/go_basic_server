package Helloworld

import (
	"encoding/json"
	"net/http"
)

func helloworldDemo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	obj := make(map[string]interface{})
	obj["message"] = "Hello world"
	_ = json.NewEncoder(w).Encode(obj)
}
