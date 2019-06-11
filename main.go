package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.handle)
	http.ListenAndServe(":5050", cors.Default().Handler(mux))

	// http://localhost:5050/ping
}
