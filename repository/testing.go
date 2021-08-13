package repository

import (
	"testing"

	"github.com/AlimByWater/coinbase_ws/config"
)

func TestDB(t *testing.T, cfg *config.Config) (*MySql, func()) {
	t.Helper()

	conn, err := Connect(cfg)
	if err != nil {
		t.Fatalf("failed to create a db instance: %v", err)
	}

	db := MySql{DB: conn}
	return &db, func() {
		db.Close()
	}
}
