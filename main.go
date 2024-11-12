package main

import (
	"log"
	"net/http"
)

type Server struct {
	address string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			http.Error(w, "404 Page Not found", http.StatusNotFound)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {

	s := &Server{address: ":8080"}

	if err := http.ListenAndServe(s.address, s); err != nil {
		log.Fatal(err)
	}
}
