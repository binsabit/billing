package main

import (
	"gitlab.com/binsabit/billing/app"
	"gitlab.com/binsabit/billing/config"
	"gitlab.com/binsabit/billing/database"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Printf("could not connect to database: %v", err)
		os.Exit(0)
	}
	defer db.Close()
	storage := database.NewStorage(db)

	server, err := app.NewServer(cfg, storage)

	if err != nil {
		log.Printf("could not create server: %v", err)
		os.Exit(0)
	}

	if err := server.Start(":" + cfg.ServerPort); err != nil {
		log.Printf("could not start router: %v ", err)
	}
}
