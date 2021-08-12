package main

import (
	"flag"
	"log"

	"github.com/AlimByWater/coinbase_ws/config"
	"github.com/AlimByWater/coinbase_ws/internal/parser"
	"github.com/AlimByWater/coinbase_ws/repository"
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

	// Connect to the db and remember to close it
	db, err := repository.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to create a db instance: %v", err)
	}
	defer db.Close()

	wsConn, err := parser.GetWSConnection()
	if err != nil {
		log.Fatalf("failed to get ws connection: %v", err)
	}
	parser := parser.NewParser(wsConn)
	defer parser.CloseWS()

	if err := parser.Subscribe(cfg.Symbols); err != nil {
		log.Fatalf("%v", err)
	}

	if err := parser.PrintLog(); err != nil {
		log.Fatalf("%v", err)
	}

}
