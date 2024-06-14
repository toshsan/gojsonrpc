package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type LoggrMux struct {
	handler http.Handler
}

func (l *LoggrMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	l.handler.ServeHTTP(w, req)
	log.Print(req.UserAgent(), req.RemoteAddr, req.Method, req.RequestURI)
}
