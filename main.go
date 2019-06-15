package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/skratchdot/open-golang/open"
)

func main() {

	go func() {
		open.Run("http://localhost:8080/")
		http.ListenAndServe(":8080", newStaticHandler())
	}()

	http.ListenAndServe(":5050", newHandler())
	// http://localhost:5050/ping?since=20190611JST&before=20190612JST"
}

func newHandler() http.Handler {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.handle)
	return cors.Default().Handler(mux)
}

func newStaticHandler() http.Handler {
	static := "./static"
	return http.FileServer(http.Dir(static))
}
