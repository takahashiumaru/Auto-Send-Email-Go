package main

import (
	api "auto-emails/api"
	"net/http"
)

// Import Gorilla Mux router

func main() {
	var w http.ResponseWriter
	var r *http.Request
	api.Handler(w, r)
}
