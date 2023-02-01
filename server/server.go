package server

import (
	"net/http"
	"fmt"
)

// ServeMux creates a new HTTP server.
func ServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)

	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("'/' is accessed")
	fmt.Fprintln(w, "Hello")
}