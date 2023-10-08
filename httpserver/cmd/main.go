package main

import (
	"fmt"
	"httpserver/internal/config"
	"httpserver/internal/server"
	"log"
)

func main() {
	fmt.Println("Server is starting...")

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(cfg)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started!")
}
