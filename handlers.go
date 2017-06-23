package main

import (
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endpoint not valid", http.StatusMethodNotAllowed)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("I feel happy!\n"))
}
