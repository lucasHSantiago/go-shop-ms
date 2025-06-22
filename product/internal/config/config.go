package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/conf/v3"
)

type Config struct {
	Web struct {
		ReadTimeout     time.Duration `conf:"default:5s"`
		WriteTimeout    time.Duration `conf:"default:10s"`
		IdleTimeout     time.Duration `conf:"default:120s"`
		ShutdownTimeout time.Duration `conf:"default:20s"`
		APIHost         string        `conf:"default:0.0.0.0:5000"`
	}
	DB struct {
		User         string `conf:"default:postgres"`
		Password     string `conf:"default:admin,mask"`
		Host         string `conf:"default:localhost:5432"`
		Name         string `conf:"default:product"`
		MaxIdleConns int    `conf:"default:2"`
		MaxOpenConns int    `conf:"default:0"`
		DisableTLS   bool   `conf:"default:true"`
	}
}

func Parse() (Config, error) {
	var cfg Config
	const prefix = "PRODUCT"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}

	return cfg, nil
}
