package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Amitaarav/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println(os.Args)
	if(len(os.Args) < 2){
		log.Fatalf("usage: migrate <up | down>")
	}

	cfg := config.MustLoad()
	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl,
	)

	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatalf("migration.up: %v", err)
		}
		log.Println("up called")
	
	case "down":
		if err := m.Down(); err != nil {
			log.Fatalf("migration.down: %v", err)
		}
		log.Printf("down called")

	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("runing migration")
}