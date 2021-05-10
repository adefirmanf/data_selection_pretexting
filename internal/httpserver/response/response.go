package response

import (
	"encoding/json"
	"net/http"
)

// Error .
type Error struct {
	Error string `json:"error"`
}

// TransformSuccessAsJSON .
func TransformSuccessAsJSON(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(data)
	return
}

// TransformErrorAsJSON .
func TransformErrorAsJSON(w http.ResponseWriter, statuscode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	resultsError := &Error{
		Error: http.StatusText(statuscode),
	}
	json.NewEncoder(w).Encode(resultsError)
	return
}
