package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main(){
	mux := http.NewServeMux()
	// why pointer ?
	mux.HandleFunc("GET /healthz", func(writer http.ResponseWriter, response *http.Request) {
		writer.Header().Set("Content-Type", "application/json") // 1. first in memory set
		writer.WriteHeader(http.StatusOK) // 2. 
		writer.Write([]byte(`{"status" : "all ok"}`)) // 3.
	})

	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}