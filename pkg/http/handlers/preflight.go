package handlers

import (
	"net/http"
)

// Preflight ...
func Preflight(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	return
}
