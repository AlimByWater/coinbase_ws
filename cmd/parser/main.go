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

	wsConn, err := parser.GetWsConnection()
	if err != nil {
		log.Fatalf("failed to get ws connection: %v", err)
	}

	parser := parser.Parser{
		WsConn: wsConn,
		DB:     &repository.MySql{DB: db},
	}
	defer parser.CloseWsConnection()

	if err := parser.Subscribe(cfg.Symbols); err != nil {
		log.Fatalf("%v", err)
	}

	if err := parser.Start(); err != nil {
		log.Fatalf("%v", err)
	}

}
