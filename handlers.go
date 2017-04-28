package main

import (
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endpoint not valid", http.StatusMethodNotAllowed)
}
