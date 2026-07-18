package db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	return "host=localhost port=5434 user=postgres password=postgres dbname=mockmaster sslmode=disable"
}

func NewConnectionString() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbConn, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
