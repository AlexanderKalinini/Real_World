package middleware

import (
	"errors"
	"net/http"
	appErrors "rwa/internal/errors"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func ErrorHandler(handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			return
		}

		status := http.StatusInternalServerError

		switch {
		case errors.As(err, &appErrors.ValidationErrors{}):
			status = http.StatusUnprocessableEntity
		case errors.As(err, &appErrors.NotFoundError{}):
			status = http.StatusNotFound
		case errors.As(err, &appErrors.UnauthorizedError{}):
			status = http.StatusUnauthorized
		}

		http.Error(w, http.StatusText(status), status)
	}
}
