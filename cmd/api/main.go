package main

import (
	"flag"
	"forum/internal/app/server"
	"log"
	"os"
)

var configWay string

func init() {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatal("Error: var JWT_KEY is not set")
	}
	flag.StringVar(&configWay, "config-way", "configs/config.json", "way to config")

}

func main() {
	flag.Parse()

	config := server.NewConfig()
	err := config.ReadConfig(configWay)
	if err != nil {
		log.Fatal("Error: can't read config file: %s\n", err)
	}

	log.Fatal(server.Start(config))
}
