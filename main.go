package main

import (
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello World"}`))
}

func main() {
	s := Server{}
	http.HandleFunc("/", s.Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
