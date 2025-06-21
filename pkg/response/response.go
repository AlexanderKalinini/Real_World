package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, v interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
