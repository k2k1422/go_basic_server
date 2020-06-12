package Response

import (
	"encoding/json"
	"net/http"
	"server/DataModels"
)

func getResponseBody(code string, data ...interface{}) DataModels.Response {
	if len(data) == 0 {
		return DataModels.Response{
			ResponseCode:    code,
			ResponseMessage: Status[code],
		}
	} else {
		return DataModels.Response{
			ResponseCode:    code,
			ResponseMessage: Status[code],
			Data:            data[0],
		}
	}

}

func BadRequest(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func Conflict(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusConflict)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func InternalServerError(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func Success(w http.ResponseWriter, r *http.Request, code string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(getResponseBody(code, data))
}

func Created(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func Unauthorized(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func Unprocessed(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func AccessDenied(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}

func ResourceUpdate(w http.ResponseWriter, r *http.Request, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(getResponseBody(code))
}
