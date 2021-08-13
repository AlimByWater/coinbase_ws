package parser

import (
	"log"
	"strconv"
	"time"

	"github.com/AlimByWater/coinbase_ws/models"
	"github.com/AlimByWater/coinbase_ws/proto"
	"github.com/pkg/errors"
	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

func (p *Parser) InsertNewTick(msg *coinbasepro.Message, errCh chan error) {
	tick, err := parseTick(msg)
	if err != nil {
		errCh <- err
		return
	}

	_, err = p.DB.InsertNewTick(tick)
	if err != nil {
		errCh <- err
	}

	log.Printf("New Tick Inserted: %+v", tick)
}

func parseTick(msg *coinbasepro.Message) (*models.Tick, error) {
	bid, err := strconv.ParseFloat(msg.BestBid, 64)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing bid: "+msg.Message)
	}

	ask, err := strconv.ParseFloat(msg.BestAsk, 64)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing ask: "+msg.Message)
	}

	return &models.Tick{
		Symbol:    proto.Symbol(msg.ProductId),
		Bid:       bid,
		Ask:       ask,
		CreatedAt: time.Now().String(),
	}, nil
}
