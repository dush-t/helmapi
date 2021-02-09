package main

import (
	"log"
	"net/http"

	"github.com/dush-t/helmapi/api"
)

func main() {

	// Routes for charts
	http.Handle("/install", api.InstallChartHandler())
	http.Handle("/delete", api.DeleteReleaseHandler())

	// Routes for repos
	http.Handle("/repo/add", api.AddRepoHandler())
	http.Handle("/repo/delete", api.RemoveRepoHandler())
	http.Handle("/repo/update", api.RepoUpdateHandler())

	// Health check endpoint
	http.Handle("/healthcheck", api.HealthCheckHandler())

	log.Println("HTTP server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
