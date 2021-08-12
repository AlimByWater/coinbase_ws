package models

import (
	"github.com/AlimByWater/coinbase_ws/proto"
	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

type Tick struct {
	Symbol    proto.Symbol `json:"symbol"`
	Ask       float64      `json:"ask"`
	Bid       float64      `json:"bid"`
	CreatedAt string       `json:"created_at"`
}

func (t *Tick) Adapt(msg *coinbasepro.Message) {
}
