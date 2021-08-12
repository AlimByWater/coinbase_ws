package parser

import (
	"log"

	"github.com/AlimByWater/coinbase_ws/config"
	"github.com/AlimByWater/coinbase_ws/repository"
	ws "github.com/gorilla/websocket"
	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

type Parser struct {
	wsConn *ws.Conn
	DB     repository.Repository
}

func NewParser(wsConn *ws.Conn) *Parser {
	return &Parser{
		wsConn: wsConn,
	}
}

func GetWSConnection() (*ws.Conn, error) {
	var d ws.Dialer
	wsConn, _, err := d.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		return nil, err
	}

	return wsConn, nil
}

func (p *Parser) Subscribe(symbols config.Symbols) error {
	subscribe := coinbasepro.Message{
		Type: "subscribe",
		Channels: []coinbasepro.MessageChannel{
			{
				Name:       "ticker",
				ProductIds: symbols.GetAllSymbols(),
			},
		},
	}

	err := p.wsConn.WriteJSON(subscribe)
	return err
}

func (p *Parser) PrintLog() error {
	for {
		message := coinbasepro.Message{}
		if err := p.wsConn.ReadJSON(&message); err != nil {
			return err
		}

		log.Printf("%s: BID %s", message.ProductId, message.BestBid)
	}
}

func (p *Parser) CloseWS() error {
	return p.wsConn.Close()
}
