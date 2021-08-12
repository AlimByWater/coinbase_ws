package repository

import (
	"context"

	"github.com/AlimByWater/coinbase_ws/models"
)

type TicksRepository interface {
	InsertNewTick(context.Context, models.Tick)
}
