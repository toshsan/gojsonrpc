package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	addr := GetEnv("ADDR", "localhost:3000")

	mux := http.NewServeMux()
	mux.Handle("/api", &JSONServer{&Foo{}})

	log.Info().Msg("Starting server at http://" + addr)
	log.Panic().Err(http.ListenAndServe(addr, &LoggrMux{mux}))
}
