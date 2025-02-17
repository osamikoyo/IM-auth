package data

import (
	"github.com/osamikoyo/IM-auth/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(cfg *config.Config) (*Storage, error) {
	g, err := gorm.Open(sqlite.Open(cfg.DSN))
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: g,
	}, nil
}


