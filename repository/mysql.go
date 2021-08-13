package repository

import (
	"fmt"
	"time"

	"github.com/AlimByWater/coinbase_ws/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	queryTimeout = 30 * time.Second
)

type Repository interface {
	TicksRepository
}

type MySql struct {
	DB *sqlx.DB
}

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", GetConnectionString(*cfg))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(queryTimeout)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

// GetPgConnectionOptions is for retriving pg.Options from configs
func GetConnectionString(cfg config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Db)

}

// GetConnection ...
func (m *MySql) GetConnection() *sqlx.DB {
	return m.DB
}

// Close ...
func (m *MySql) Close() error {
	return m.DB.Close()
}
