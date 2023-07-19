package main

import (
	"gitlab.com/binsabit/billing/app"
	"gitlab.com/binsabit/billing/config"
	"gitlab.com/binsabit/billing/db"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := db.Connect(cfg)
	if err != nil {
		log.Printf("could not connect to models: %v", err)
		os.Exit(0)
	}
	defer conn.Close()
	storage := db.NewStore(conn)

	server, err := app.NewServer(cfg, storage)

	if err != nil {
		log.Printf("could not create server: %v", err)
		os.Exit(0)
	}

	if err := server.Start(":" + cfg.ServerPort); err != nil {
		log.Printf("could not start router: %v ", err)
	}
}
