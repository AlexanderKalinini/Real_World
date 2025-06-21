package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	appErrors "rwa/internal/errors"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func ErrorHandler(handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		var status = http.StatusInternalServerError

		var ve *appErrors.ValidationErrors
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		if errors.As(err, &ve) {
			status = http.StatusUnprocessableEntity
		}

		if errors.As(err, &syntaxError) {
			status = http.StatusUnprocessableEntity
		}

		if errors.As(err, &unmarshalTypeError) {
			status = http.StatusUnprocessableEntity
			err = fmt.Errorf("Неверный тип данных в поле %s. Ожидается тип: %s ", unmarshalTypeError.Field, unmarshalTypeError.Type)
		}

		var nfe *appErrors.NotFoundError
		if errors.As(err, &nfe) {
			status = http.StatusNotFound
		}

		var ue *appErrors.UnauthorizedError
		if errors.As(err, &ue) {
			status = http.StatusUnauthorized
		}
		errorData := map[string]map[string][]string{
			"errors": {
				"body": {
					err.Error(),
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(errorData)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
