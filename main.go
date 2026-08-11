package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main(){
	mux := http.NewServeMux()
	// why pointer ?
	mux.HandleFunc("GET /healthz", func(writer http.ResponseWriter, response *http.Request) {
		writer.Header().Set("Content-Type", "application/json") // 1. first in memory set
		writer.WriteHeader(http.StatusOK) // 2. 
		writer.Write([]byte(`{"status" : "all ok"}`)) // 3.
	})

	srv := http.Server{
		Addr: ":8090",
		Handler: mux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second *30,
		IdleTimeout: time.Second * 60,
	}
	err := srv.ListenAndServe() // ReadTimeout, WriteTimeout, IdleTimeout as default are zero , to make it have industry relevant timeouts use srv
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}