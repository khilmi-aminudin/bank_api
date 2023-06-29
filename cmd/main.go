package main

import (
	"log"

	"github.com/khilmi-aminudin/bank_api/server"
	"github.com/khilmi-aminudin/bank_api/utils"
)

func main() {
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server, err := server.NewServer(*cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	log.Fatal(server.Start(cfg.ServerAddress))
}
