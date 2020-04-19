package handlers

import (
	"net/http"
	"time"

	"github.com/da-moon/coe865-final/pkg/http/response"
)

// HealthCheck ...
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	response.WriteSuccessfulJSON(&w, r, map[string]interface{}{
		"status": "healthy",
		"time":   time.Now().Unix(),
	})
	return
}
