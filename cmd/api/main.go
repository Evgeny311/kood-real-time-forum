package main

import (
	"flag"
	"kood-real-time-forum/internal/app/server"
	"log"
	"os"
)

var configPath string

func init() {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatalf("error: env variable JWT_KEY is not set")
	}

	flag.StringVar(&configPath, "config-path", "configs/config.json", "path to config")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	log.Fatal(server.Start(config))
}
