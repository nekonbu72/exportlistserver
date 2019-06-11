package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	http.ListenAndServe(":5050", newHandler())

	// http://localhost:5050/ping
}

func newHandler() http.Handler {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.handle)
	return cors.Default().Handler(mux)
}
