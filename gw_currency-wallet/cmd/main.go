package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/emmonbear/wallet-exchanger/internal/config"
)

func main() {
	configPath := flag.String("c", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		log.Fatalf("Config file path not provided. Use -c flag to specify the path")
	}

	cfg := config.MustLoad(*configPath)
	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: postgres

	// TODO: init router: chi, "chi render" (gin)

	// TODO: run server:
}
