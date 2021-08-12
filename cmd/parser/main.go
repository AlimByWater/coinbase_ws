package main

import (
	"flag"
	"log"

	"github.com/AlimByWater/coinbase_ws/config"
)

const (
	defaultConfigPath = "config/parser.toml"
	defaultSymbols    = "ETH-BTC,BTC-USD,BTC-EUR"
)

func main() {
	// Parse flags
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	symbolsFlag := flag.String("symbols", defaultSymbols, "symbols to parse in format like ETH-BTC,BTC-USD")
	flag.Parse()

	cfg, err := config.Parse(*configPath, *symbolsFlag)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	cfg.Print() // just for debugging
}
