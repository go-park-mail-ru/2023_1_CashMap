package server

import (
	"net/http"
	"strconv"
)

type Server struct {
	*http.Server
}

func NewServer(handler http.Handler, port uint) *Server {
	return &Server{
		&http.Server{
			Addr:    ":" + strconv.FormatUint(uint64(port), 10),
			Handler: handler,
		},
	}
}
