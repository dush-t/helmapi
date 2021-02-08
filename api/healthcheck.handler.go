package api

import (
	"encoding/json"
	"net/http"
)

// HealthCheckHandler serves requests at /healthcheck
func HealthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		payload := struct {
			Status string `json:"status"`
		}{Status: "OK"}
		json.NewEncoder(w).Encode(payload)
	})
}
