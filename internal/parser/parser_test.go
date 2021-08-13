package parser_test

import (
	"log"
	"os"
	"testing"

	"github.com/AlimByWater/coinbase_ws/config"
	"github.com/AlimByWater/coinbase_ws/internal/parser"
	ws "github.com/gorilla/websocket"
	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

var (
	testCfg *config.Config
)

func TestMain(m *testing.M) {
	var err error
	testCfg, err = config.Parse("../../config/parser.toml", "ETH-BTC,BTC-USD")
	if err != nil {
		log.Fatalf("%v", err)
	}

	os.Exit(m.Run())
}

func TestParser_Listen(t *testing.T) {
	parser, close := parser.TestParser(t, testCfg)
	defer close()

	if err := parser.Subscribe(testCfg.Symbols); err != nil {
		t.Fatalf("%v", err)
	}

	errCh := make(chan error)
	msgCh := make(chan *coinbasepro.Message)

	for {
		select {
		case err := <-errCh:
			t.Fatal(err)
		case msg := <-msgCh:
			t.Log(msg)
		default:
			go parser.ReadNext(msgCh, errCh)
		}
	}

}

func TestParser_GetConnection(t *testing.T) {
	var d ws.Dialer
	wsConn, resp, err := d.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)

	subscribe := coinbasepro.Message{
		Type: "subscribe",
		Channels: []coinbasepro.MessageChannel{
			{
				Name:       "ticker",
				ProductIds: testCfg.Symbols.GetAllSymbols(),
			},
		},
	}

	err = wsConn.WriteJSON(subscribe)
	if err != nil {
		t.Fatal(err)
	}

	message := coinbasepro.Message{}
	if err := wsConn.ReadJSON(&message); err != nil {
		t.Fatal(err)
	}

	for {
		if message.ProductId == "" || message.Type == "" || message.Channels == nil {
			t.Fatal("FFAT")
		}

		t.Logf("%s, BID: %s", message.ProductId, message.BestBid)
	}

}
