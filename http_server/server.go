package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

func NewServer(options ...func(*Server)) *Server {
	s := &Server{mux: http.NewServeMux()}

	for _, f := range options {
		f(s)
	}

	if s.logger == nil {
		s.logger = log.New(os.Stdout, "", 0)
	}

	s.mux.HandleFunc("/", s.index)

	return s
}

func Logger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	time.Sleep(TIMEOUT - time.Second)
	w.Write([]byte("Hello world!"))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "example Go server")

	s.mux.ServeHTTP(w, r)
}
