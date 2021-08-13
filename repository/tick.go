package repository

import (
	"github.com/AlimByWater/coinbase_ws/models"
	"github.com/pkg/errors"
)

type TicksRepository interface {
	InsertNewTick(*models.Tick) (int64, error)
}

func (m *MySql) InsertNewTick(tick *models.Tick) (int64, error) {
	res, err := m.DB.Exec(
		"INSERT INTO tick (symbol, bid, ask, created_at) VALUES (?, ?, ?, ?)",
		tick.Symbol,
		tick.Bid,
		tick.Ask,
		tick.CreatedAt,
	)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert ticket")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
