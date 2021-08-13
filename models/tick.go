package models

import (
	"github.com/AlimByWater/coinbase_ws/proto"
)

type Tick struct {
	ID        uint64       `json:"id,omitempty" db:"id"`
	Symbol    proto.Symbol `json:"symbol" db:"symbol"`
	Ask       float64      `json:"ask" db:"ask"`
	Bid       float64      `json:"bid" db:"bid"`
	CreatedAt string       `json:"created_at" db:"created_at"`
}
