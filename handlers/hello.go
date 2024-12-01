package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	response := fmt.Sprintf("Hello, %s! Welcome to Go Docker Demo", name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
	}{
		Status:    "healthy",
		Timestamp: time.Now(),
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Status: %s, Time: %s", status.Status, status.Timestamp)))
}
