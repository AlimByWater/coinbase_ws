package repository_test

import (
	"log"
	"os"
	"testing"

	"github.com/AlimByWater/coinbase_ws/config"
)

var (
	testCfg *config.Config
)

func TestMain(m *testing.M) {
	var err error
	testCfg, err = config.Parse("../config/parser.toml", "ETH-BTC")
	if err != nil {
		log.Fatalf("%v", err)
	}

	os.Exit(m.Run())
}
