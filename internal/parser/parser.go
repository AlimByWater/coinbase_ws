package parser

import (
	"github.com/AlimByWater/coinbase_ws/config"
	"github.com/AlimByWater/coinbase_ws/repository"
	ws "github.com/gorilla/websocket"
	"github.com/pkg/errors"
	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

type Parser struct {
	WsConn *ws.Conn
	DB     repository.Repository
}

// func NewParser(WsConn *ws.Conn, db *repository.Repository) *Parser {
// 	return &Parser{
// 		WsConn: WsConn,
// 		DB:     db,
// 	}
// }

func GetWsConnection() (*ws.Conn, error) {
	var d ws.Dialer
	WsConn, _, err := d.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		return nil, err
	}

	return WsConn, nil
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

	err := p.WsConn.WriteJSON(subscribe)
	return err
}

func (p *Parser) Start() error {
	errCh := make(chan error)
	msgCh := make(chan *coinbasepro.Message)

	for {
		select {
		case err := <-errCh:
			return err
		default:
			p.ReadNext(msgCh, errCh)
		}
	}
}

func (p *Parser) ReadNext(msgCh chan *coinbasepro.Message, errCh chan error) {
	message := &coinbasepro.Message{}
	if err := p.WsConn.ReadJSON(message); err != nil {
		errCh <- errors.Wrap(err, "error reading json")
		return
	}

	if message.ProductId == "" {
		return
	}

	p.InsertNewTick(message, errCh)
}

func (p *Parser) CloseWsConnection() error {
	return p.WsConn.Close()
}
