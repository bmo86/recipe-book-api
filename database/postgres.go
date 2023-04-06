package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type instancePostgres struct {
	db *gorm.DB
}

func NewConnPostgres(url string) (*instancePostgres, error) {
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		return nil, err
	}

	return &instancePostgres{db}, nil
}
