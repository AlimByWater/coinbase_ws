package config

import (
	"fmt"
	"strings"

	"github.com/AlimByWater/coinbase_ws/proto"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

type Config struct {
	Database Database
	Symbols  Symbols
}

func (c *Config) validate() error {
	return multierr.Combine(
		c.Symbols.validate(),
		c.Database.validate(),
	)
}

type Symbols []proto.Symbol

func (s *Symbols) Parse(symbolsString string) {
	symbols := strings.Split(symbolsString, ",")

	for i := range symbols {
		*s = append(*s, proto.Symbol(symbols[i]))
	}
}

func (s Symbols) validate() error {
	for i := range s {
		if err := s[i].Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Db       string
}

func (d *Database) validate() error {
	if d.Host == "" {
		return errors.New("empty db host provided")
	}
	if d.Port == 0 {
		return errors.New("empty db port provided")
	}
	if d.User == "" {
		return errors.New("empty db user provided")
	}
	if d.Password == "" {
		return errors.New("empty db password provided")
	}
	if d.Db == "" {
		return errors.New("empty db name provided")
	}
	return nil
}

func Parse(configPath, symbols string) (*Config, error) {
	setDefaults()

	// Parse the file
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}

	// Unmarshal the config
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	cfg.Symbols.Parse(symbols)

	// Validate the provided configuration
	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate the config")
	}
	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("Database.Host", "")
	viper.SetDefault("Database.Port", 0)
	viper.SetDefault("Database.User", "")
	viper.SetDefault("Database.Password", "")
	viper.SetDefault("Database.Db", "")

	viper.SetDefault("Symbols", make(Symbols, 0))
}

func (c *Config) Print() {
	inspected := *c // get a copy of an actual object
	// Hide sensitive data

	inspected.Database.User = ""
	inspected.Database.Password = ""

	fmt.Printf("%+v\n", inspected)
}
