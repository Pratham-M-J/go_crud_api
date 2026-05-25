package response

import (
	"net/http"
	"encoding/json"
)

type Response struct{
	Status string
	Error string
}

const (
	StatusOK     = "OK"
	StatusError  = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error { //interface means - data can hold a value of any type.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response{
	return Response {
		Status: StatusError,
		Error: err.Error(),
	}
}