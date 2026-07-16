package db

import (
	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	return "host=localhost port=5434 user=postgres password=postgres dbname=mockmaster"
}

func NewConnectionString() (*sqlx.DB, error) {
	dbSurce := GetConnectionString()
	dbConn, err := sqlx.Connect("postgres", dbSurce)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
