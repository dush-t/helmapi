package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dush-t/helmAPI/client"
)

// InstallChartHandler serves requests at /install
func InstallChartHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ir client.InstallRequest
		err := json.NewDecoder(r.Body).Decode(&ir)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		createErr := ir.Execute()
		if createErr != nil {
			log.Println(createErr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		payload := struct {
			Status string `json:"status"`
		}{Status: "SUCCESS"}
		json.NewEncoder(w).Encode(payload)
	})
}

// DeleteReleaseHandler serves requests at /delete
func DeleteReleaseHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dr client.DeleteRequest
		err := json.NewDecoder(r.Body).Decode(&dr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		deleteErr := dr.Execute()
		if deleteErr != nil {
			log.Println(deleteErr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		payload := struct {
			Status string `json:"status"`
		}{Status: "SUCCESS"}
		json.NewEncoder(w).Encode(payload)
	})
}
