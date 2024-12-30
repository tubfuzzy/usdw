package db

import (
	"usdw/config"
)

type DB struct {
}

func NewDB(conf *config.Configuration) (*DB, error) {

	return &DB{}, nil
}
