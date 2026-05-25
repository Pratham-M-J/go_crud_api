package student

import (
	"net/http"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"fmt"

	"github.com/Pratham-M-J/crud_api/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/Pratham-M-J/crud_api/internal/utils/response"
)

func New() http.HandlerFunc{ //inject dependency as an argument
	return func(w http.ResponseWriter, r *http.Request){

		var student types.Student
		slog.Info("Creating a student")
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){ //if the body is empty, i.e. EOF
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Empty body")))
			return
		}
		if err != nil{
			validateErrs := err.(validator)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		// validation of request
		if err:= validator.New().Struct(student); err != nil{
			response.WriteJson(w, http.StatusBadRequest, response.validationError(err))
		}

		response.WriteJson(w, http.StatusCreated, "Yo yo yo yo")
	}
}