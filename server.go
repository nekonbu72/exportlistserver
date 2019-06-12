package main

import "net/http"

type Server struct {
}

func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGet(w, r)
		return
	case "POST":
		s.handlePost(w, r)
		return
	case "DELETE":
		s.handleDelete(w, r)
		return
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		respond(w, r, http.StatusOK, nil)
		return
	}
	// not found
	respondHTTPErr(w, r, http.StatusNotFound)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	result, err := businessLogic(NewPath(r.URL.String()))
	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, err)
		return
	}
	respond(w, r, http.StatusOK, result)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
}
