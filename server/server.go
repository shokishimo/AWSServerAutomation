package server

import (
	"net/http"
)

// ServeMux creates a new HTTP server.
func ServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func IndexHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	return mux
}