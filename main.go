package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	http.ListenAndServe(":5050", newHandler())

	// http://localhost:5050/ping?since=20190611JST&before=20190612JST"
}

func newHandler() http.Handler {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.handle)
	return cors.Default().Handler(mux)
}
