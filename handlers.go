package main

import (
	"net/http"
)

type query struct {
	Entry string
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endpoint not valid", http.StatusMethodNotAllowed)
}
