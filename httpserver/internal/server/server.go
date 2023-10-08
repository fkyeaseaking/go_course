package server

import (
	"fmt"
	"httpserver/internal/config"
	"net/http"
)

func New(cfg config.Config) Server {
	s := Server{port: cfg.Port}
	http.HandleFunc("/useragent", s.handleUserAgent)
	http.HandleFunc("/myip", s.handleMyIp)

	return s
}

type Server struct {
	port int
}

func (s Server) Start() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)

	return err
}

func (s Server) handleUserAgent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, _ = w.Write([]byte("your user agent is "))
	_, _ = w.Write([]byte(r.UserAgent()))
}

func (s Server) handleMyIp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, _ = w.Write([]byte("your ip is "))
	_, _ = w.Write([]byte(r.RemoteAddr))
}
