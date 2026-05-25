package response

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

)

type Response struct{
	Status string `json:"status"`
	Error string  `json:"error"`
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

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.Tag() {

		case "required":
			errMsgs = append(errMsgs,
				fmt.Sprintf("%s is required", err.Field()),
			)

		case "email":
			errMsgs = append(errMsgs,
				fmt.Sprintf("%s must be a valid email", err.Field()),
			)

		case "gte":
			errMsgs = append(errMsgs,
				fmt.Sprintf("%s must be greater than or equal to %s",
					err.Field(),
					err.Param(),
				),
			)

		case "lte":
			errMsgs = append(errMsgs,
				fmt.Sprintf("%s must be less than or equal to %s",
					err.Field(),
					err.Param(),
				),
			)

		default:
			errMsgs = append(errMsgs,
				fmt.Sprintf("%s is invalid", err.Field()),
			)
		}
	}

	return Response{
		Status:  "Error",
		Error: strings.Join(errMsgs, ", "),
	}
}
