package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dush-t/helmAPI/client"
)

// AddRepoHandler serves requests at /repo/add
func AddRepoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ra client.RepoAddRequest
		err := json.NewDecoder(r.Body).Decode(&ra)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		addErr := ra.Execute()
		if addErr != nil {
			log.Println(addErr)
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

// RemoveRepoHandler serves requests at /repo/delete
func RemoveRepoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rr client.RepoRemoveRequest
		err := json.NewDecoder(r.Body).Decode(&rr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		removeErr := rr.Execute()
		if removeErr != nil {
			log.Println(removeErr)
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

// RepoUpdateHandler serves requests at /repo/update
func RepoUpdateHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := client.UpdateRepos(); err != nil {
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
