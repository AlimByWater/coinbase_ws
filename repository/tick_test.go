package repository_test

import (
	"testing"

	"github.com/AlimByWater/coinbase_ws/models"
	"github.com/AlimByWater/coinbase_ws/repository"
	"github.com/stretchr/testify/assert"
)

func TestTick_InsertNewTick(t *testing.T) {
	db, close := repository.TestDB(t, testCfg)
	defer close()

	tick := &models.Tick{
		Symbol: testCfg.Symbols[0],
		Bid:    15,
		Ask:    12,
	}

	i, err := db.InsertNewTick(tick)
	assert.NoError(t, err)
	t.Log(i)
}
