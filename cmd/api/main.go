package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // it adds router to global
	"time"

	"github.com/Amitaarav/olx-api/internal/config"
	"github.com/Amitaarav/olx-api/internal/db"
	"github.com/Amitaarav/olx-api/internal/handlers"
	"github.com/joho/godotenv"
)

func main(){
	cfg := config.MustLoad()

	_, dbErr := db.Connect(cfg.DatabaseUrl)
	if dbErr != nil{
		log.Fatalf("main.db.connect: %v", dbErr)
	}

	err := godotenv.Load()
	
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Printf("starting olx server...")
	mux := http.NewServeMux() // own router, not global
	// why pointer ?
	mux.HandleFunc("GET /healthz", handlers.Health)

	// HandleFunc calls DefaultServeMux(default router) and registering, which is globally allocated in the application

	srv := http.Server{
		Addr: ":" + cfg.Port,
		Handler: mux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second *30,
		IdleTimeout: time.Second * 60,
	}
	
	if err := srv.ListenAndServe(); // ReadTimeout, WriteTimeout, IdleTimeout as default are zero , to make it have industry relevant timeouts use srv

	err != nil {
		log.Fatalf("server failed: %v", err)
	}
}