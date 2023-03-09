package server

import "net/http"

type Server struct {
	*http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}
