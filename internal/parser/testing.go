package parser

import (
	"log"
	"testing"

	"github.com/AlimByWater/coinbase_ws/config"
)

func TestParser(t *testing.T, cfg *config.Config) (*Parser, func()) {

	wsConn, err := GetWsConnection()
	if err != nil {
		log.Fatalf("failed to get ws connection: %v", err)
	}

	parser := &Parser{WsConn: wsConn}
	return parser, func() {
		parser.CloseWsConnection()
	}
}
